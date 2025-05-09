package uptime

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"time"

	"go.etcd.io/bbolt"
	"go.uber.org/zap"
)

var (
	dbIntervals = []byte("intervals")
)

// BoltDBStore persists uptime interval data to a flat, on-disk BoltDB database.
type BoltDBStore struct {
	db     *bbolt.DB
	logger *zap.Logger
}

func NewBoltDBStore(logger *zap.Logger, path string, opts *bbolt.Options) (*BoltDBStore, error) {
	if opts == nil {
		opts = &bbolt.Options{}
	}

	db, err := bbolt.Open(path, 0600, opts)
	if err != nil {
		return nil, fmt.Errorf("opening database: %v", err)
	}
	store := &BoltDBStore{
		db:     db,
		logger: logger,
	}

	// If the database is writeable, ensure our buckets are ready.
	if !opts.ReadOnly {
		if err = store.initialize(); err != nil {
			store.Close()
			return nil, fmt.Errorf("initializing: %v", err)
		}
	}

	return store, nil
}

func (s *BoltDBStore) initialize() error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists(dbIntervals); err != nil {
			return err
		}
		return nil
	})
}

func (s *BoltDBStore) Close() error {
	return s.db.Close()
}

func (s *BoltDBStore) Put(_ context.Context, data IntervalData) error {
	key := keyFromTime(data.StartTime)
	val, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("encoding interval: %v", err)
	}

	return s.db.Update(func(tx *bbolt.Tx) error {
		return tx.Bucket(dbIntervals).Put(key, val)
	})
}

func (s *BoltDBStore) Range(_ context.Context, start, end time.Time, fn func(IntervalData) error) error {
	// Start a read-only transaction
	tx, err := s.db.Begin(false)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Define the limits of the range scan
	c := tx.Bucket(dbIntervals).Cursor()
	min := keyFromTime(start)
	max := keyFromTime(end)

	for k, v := c.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = c.Next() {
		var data IntervalData
		if err := json.Unmarshal(v, &data); err != nil {
			return fmt.Errorf("decoding interval: %v", err)
		}
		if data.StartTime.Before(start) || data.EndTime.After(end) {
			continue
		}
		if err := fn(data); err != nil {
			return err
		}
	}

	return nil
}

// keyFromTime packs the Unix timestamp of the provided Time into 8 bytes.
// This allows the database to order records by time with 1 second granularity.
func keyFromTime(t time.Time) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(t.Unix()))
	return buf
}
