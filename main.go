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
	serialNumber string
	isUSB bool
	list bool
)

func main() {
	flag.StringVar(&vid, "vid", "", "VID, default null")
	flag.StringVar(&pid, "pid", "", "PID, default null")
	flag.StringVar(&serialNumber, "serialnumber", "", "Serial Number, default null")
	flag.BoolVar(&isUSB, "usb", true, "USB. Default true")
	flag.BoolVar(&list, "list", false, "List Ports. Default false")
	flag.Parse()

	// if isUSB && (vid == nil || len(*vid) < 1 || pid == nil || len(*pid) < 1) {
	// 	log.Fatal("Incomplete argrument")
	// }

	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		log.Fatal(err)
	}
	
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
		return
	}

	if list {
		for _, port := range ports {
			fmt.Printf("Found port: %s\n", port.Name)
			if port.IsUSB {
				fmt.Printf("   USB VID    %s\n", port.VID)
				fmt.Printf("   USB PID    %s\n", port.PID)
				fmt.Printf("   USB prod.  %s\n", port.Product)
				fmt.Printf("   USB serial %s\n", port.SerialNumber)
				fmt.Printf("Detailed struct: %#v\n\n", p)
			}
		}
	} else {
		for _, port := range ports {
			if port.IsUSB == isUSB && ( vid != "" || port.VID == vid ) && ( pid != "" || port.PID == pid ) && ( serialNumber != "" || port.SerialNumber == serialNumber )  {
				fmt.Print(port.Name)
			}
		}
	}
}
