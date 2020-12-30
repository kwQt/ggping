package main

import (
	"flag"
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/go-ping/ping"
	"github.com/kwQt/ggping"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "invalid arguments")
		os.Exit(1)
	}
	arg := flag.Args()
	addr := arg[0]

	err := ui.Init()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer ui.Close()

	width, height := ui.TerminalDimensions()
	status := ggping.NewStatus(width)
	plot := widgets.NewPlot()
	event := ui.PollEvents()

	pinger, err := ping.NewPinger(addr)
	if err != nil {
		ui.Close()
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	done := make(chan struct{})
	result := make(chan ggping.Result)
	go ggping.PingLoop(pinger, result, done)

loop:
	for {
		select {
		case e := <-event:
			switch e.ID {
			case "q", "<C-c>":
				pinger.Stop()
				break loop
			}
		case r := <-result:
			if err = r.Error; err != nil {
				pinger.Stop()
				ui.Close()
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			status.Update(float64(r.Packet.Rtt.Milliseconds()))
			ggping.DrawChart(plot, status.GetAll(), status.GetMax(), width, height)
		}
	}
	<-done
}
