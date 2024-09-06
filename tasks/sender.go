package tasks

import (
	"errors"
	"fmt"
	"net"
	"sync"
	"time"
)

type Sender struct {
	addr              *net.UDPAddr
	manager           *CopyManager
	heartbeatInterval time.Duration
}

func NewSender(
	addr *net.UDPAddr,
	manager *CopyManager,
	heartbeatInterval time.Duration,
) (*Sender, error) {
	if addr == nil || !addr.IP.IsMulticast() {
		return nil, errors.New("require multicast address")
	}
	if manager == nil {
		return nil, errors.New("require copy manager")
	}

	return &Sender{
		addr,
		manager,
		heartbeatInterval,
	}, nil
}

func (s *Sender) Start(group *sync.WaitGroup) {
	conn, err := net.DialUDP("udp", nil, s.addr)
	if err != nil {
		panic(err)
	}
	defer group.Done()
	defer func(conn *net.UDPConn) {
		_ = conn.Close()
	}(conn)
	fmt.Println("sender work")

	for {
		_, err := conn.Write([]byte{})
		if err != nil {
			panic(err)
		}

		time.Sleep(s.heartbeatInterval)
	}
}
