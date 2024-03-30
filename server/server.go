package server

import "time"

type Server struct {
	TotalConnections int16 `json:"totalConnections"`
	Weight  int
	responseTime time.Duration
	
}
