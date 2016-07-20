package main

import (
	"fmt"
	"strconv"
	"os"
	"time"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

func main() {
	interval := int64(30)
	s := os.Getenv("DYNO_NETWORK_STATS_INTERVAL")
	if s != "" {
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil || i < 1 {
			fmt.Println("Ignoring DYNO_NETWORK_STATS_INTERVAL value:", s)
		} else {
			interval = i
		}
	}

	source := os.Getenv("DYNO")


	// Stats are differential, so this is the "previous" reading
	prev, err := linuxproc.ReadNetworkStat("/proc/net/dev")
	if err != nil {
		panic(fmt.Sprintf("network stat read fail", err))
	}

	for {
		time.Sleep(time.Duration(interval) * time.Second)

		stat, err := linuxproc.ReadNetworkStat("/proc/net/dev")
		if err != nil {
			fmt.Println("network stat read fail", err)
		}

		for _, s := range stat {
			// Find the previous reading, or default to blank if we haven't seen this interface before
			p := linuxproc.NetworkStat{s.Iface, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
			for _, x := range prev {
				if x.Iface == s.Iface {
					p = x
				}
			}

			fmt.Printf("ps=%s sample#iface=%s sample#bytes_recv=%d sample#bytes_sent=%d sample#packets_recv=%d sample#packets_sent=%d sample#recv_error=%d sample#sent_error=%d sample#recv_drop=%d sample#sent_drop=%d\n",
				source,
				s.Iface,
				s.RxBytes - p.RxBytes,
				s.TxBytes - p.TxBytes,
				s.RxPackets - p.RxPackets,
				s.TxPackets - p.TxPackets,
				s.RxErrs - p.RxErrs,
				s.TxErrs - p.TxErrs,
				s.RxDrop - p.RxDrop,
				s.TxDrop - p.TxDrop)
		}

		prev = stat
	}
}
