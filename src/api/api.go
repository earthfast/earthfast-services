package api

type UptimeResponse struct {
	StartTime    int                          `json:"startTime"`
	EndTime      int                          `json:"endTime"`
	RequestCount int                          `json:"requestCount"`
	ProbeResults map[string]UptimeProbeCounts `json:"probeResults"`
}

type UptimeProbeCounts struct {
	Success int `json:"success"`
	Failure int `json:"failure"`
}
