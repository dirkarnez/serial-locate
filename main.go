package main

import (
	"flag"
	"fmt"
	"log"

	"go.bug.st/serial/enumerator"
)

var (
	vid   string
	pid   string
	isUSB bool
)

func main() {
	flag.StringVar(&vid, "vid", "", "VID")
	flag.StringVar(&pid, "pid", "", "PID")
	flag.BoolVar(&isUSB, "usb", true, "USB. Default true")
	flag.Parse()

	if isUSB && (len(vid) < 1 || len(pid) < 1) {
		log.Fatal("Incomplete argrument")
	}

	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		return
	}

	for _, port := range ports {
		if port.IsUSB == isUSB && port.VID == vid && port.PID == pid {
			fmt.Print(port.Name)
		}
	}
}
