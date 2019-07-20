package main

import "time"

// Client Struct to store Client Information
type Client struct {
	CommonName     string
	RealAddress    string
	BytesReceived  string
	BytesSent      string
	ConnectedSince time.Time
}

// Routing Struct to store Routing Information
type Routing struct {
	VirtualAddress string
	CommonName     string
	RealAddress    string
	LastRef        time.Time
}

// GlobalStats struct to store Global stats
type GlobalStats struct {
	MaxBcastMcastQueueLen int
}

// Status struct to store the status
type Status struct {
	ClientList   []Client
	RoutingTable []Routing
	GlobalStats  GlobalStats
	UpdatedAt    time.Time
	IsUp         bool
}

type parseError struct {
	s string
}

func (e *parseError) Error() string {
	return e.s
}
