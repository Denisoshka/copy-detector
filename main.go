package main

import (
	"flag"
	"fmt"
	"lab1/helpers"
	"lab1/tasks"
	"net"
	"strconv"
	"sync"
	"time"
)

var (
	heartbeatInterval = 2 * time.Second
)

func main() {
	host, port, err := helpers.ParseFlags()
	if err != nil {
		fmt.Println(err)
		flag.PrintDefaults()
		return
	}
	target := net.JoinHostPort(host, strconv.Itoa(port))
	addr, err := net.ResolveUDPAddr(
		"udp",
		target,
	)
	manager := tasks.NewCopyManager(heartbeatInterval)
	receiver, err := tasks.NewReceiver(addr, manager)
	if err != nil {
		fmt.Println(err)
		return
	}
	sender, err := tasks.NewSender(addr, manager, heartbeatInterval/2)
	go manager.Start()
	var group sync.WaitGroup
	group.Add(2)
	go sender.Start(&group)
	go receiver.Start(&group)
	group.Wait()
}
