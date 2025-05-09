package content

import (
	"net/http"
	"strconv"
	"time"

	"armada-node/api"
	"armada-node/metering/uptime"

	"go.uber.org/zap"
)

func (h *Handler) uptimeHandler(w http.ResponseWriter, r *http.Request) {
	start, err := strconv.ParseInt(r.URL.Query().Get("start"), 10, 32)
	if err != nil {
		http.Error(w, "Invalid query parameter: start", http.StatusBadRequest)
		return
	}
	startTime := time.Unix(start, 0)

	end, err := strconv.ParseInt(r.URL.Query().Get("end"), 10, 32)
	if err != nil {
		http.Error(w, "Invalid query parameter: end", http.StatusBadRequest)
		return
	}
	endTime := time.Unix(end, 0)

	resp := api.UptimeResponse{
		StartTime:    int(start),
		EndTime:      int(end),
		ProbeResults: make(map[string]api.UptimeProbeCounts),
	}
	err = h.uptime.Results(r.Context(), startTime, endTime, func(data uptime.IntervalData) error {
		resp.RequestCount += int(data.RequestCount)

		for host, counts := range data.ProbeResults {
			aggCounts := resp.ProbeResults[host]
			aggCounts.Success += int(counts.Success)
			aggCounts.Failure += int(counts.Failure)
			resp.ProbeResults[host] = aggCounts
		}

		return nil
	})
	if err != nil {
		h.logger.Error("Failed to fetch uptime results", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	h.sendJSON(resp, w, r)
}
