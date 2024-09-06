package tasks

import (
	"lab1/helpers"
	"sync"
	"time"
)

type CopyManager struct {
	data                map[string]time.Time
	expiredCheckTimeout time.Duration
	mutex               sync.Mutex
}

func NewCopyManager(expiredCheckTimeout time.Duration) *CopyManager {
	return &CopyManager{
		make(map[string]time.Time),
		expiredCheckTimeout,
		sync.Mutex{},
	}
}

func (cm *CopyManager) update(copy string) {
	defer cm.mutex.Unlock()
	cm.mutex.Lock()
	cm.data[copy] = time.Now().UTC()
}

func (cm *CopyManager) remove(copy string) {
	delete(cm.data, copy)
}

func (cm *CopyManager) Start() {
	for {
		cm.mutex.Lock()
		for addr, lastSeen := range cm.data {
			if time.Since(lastSeen) > cm.expiredCheckTimeout {
				cm.remove(addr)
			}
		}
		helpers.UpdateConsole(cm.data)
		cm.mutex.Unlock()
		time.Sleep(cm.expiredCheckTimeout)
	}
}
