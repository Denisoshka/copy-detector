package helpers

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"time"
)

func ParseFlags() (host string, port int, err error) {
	flag.StringVar(&host, "address", "", "Multicast group address")
	flag.IntVar(&port, "port", -1, "Multicast group port")
	flag.Parse()

	portCheck := isCorrectPort(port)
	if portCheck != nil {
		return "", 0, portCheck
	}
	addrCheck := isCorrectAddress(host)
	if addrCheck != nil {
		return "", 0, addrCheck
	}
	return host, port, nil
}

func UpdateConsole(data map[string]time.Time) {
	fmt.Print("\033[H\033[2J")
	if len(data) != 0 {
		for ip := range data {
			fmt.Println(ip)
		}
	}
}

func isCorrectAddress(addr string) error {
	_, err := net.ResolveIPAddr("ip", addr)
	if err != nil {
		return err
	}
	return nil
}

func isCorrectPort(port int) error {
	if port == -1 {
		return errors.New("port not specified")
	}
	if !(0 <= port && port <= 65536) {
		return errors.New("incorrect port")
	}
	return nil
}
