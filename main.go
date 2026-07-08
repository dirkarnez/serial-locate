package main

import (
	"flag"
	"fmt"
	"log"

	"go.bug.st/serial/enumerator"
)

var (
	vid   *string
	pid   *string
	serialNumber *string
	configuration *string
	isUSB bool
)

func main() {
	flag.StringVar(vid, "vid", nil, "VID, default null")
	flag.StringVar(pid, "pid", nil, "PID, default null")
	flag.StringVar(serialNumber, "serialnumber", nil, "Serialnumber, default null")
	flag.StringVar(configuration, "configuration", nil, "Configuration, default null")
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
		log.Fatal("No serial ports found!")
		return
	}

	for _, port := range ports {
		if port.IsUSB == isUSB && ( vid != nil || port.VID == *vid ) && ( pid != nil || port.PID == *pid )  && ( serialNumber != nil || port.SerialNumber == *serialNumber ) && ( configuration != nil || port.Configuration == *configuration )  {
			fmt.Print(port.Name)
		}
	}
}
