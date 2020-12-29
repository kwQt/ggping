package ggping

import (
	"github.com/go-ping/ping"
	"time"
)

type Result struct {
	Packet *ping.Packet
	Error  error
}

func PingLoop(pinger *ping.Pinger, result chan Result, done chan struct{}) {
	defer close(done)

	pinger.OnRecv = func(packet *ping.Packet) {
		result <- Result{packet, nil}
	}

	pinger.Interval = 100 * time.Millisecond

	err := pinger.Run()
	if err != nil {
		result <- Result{nil, err}
	}
}
