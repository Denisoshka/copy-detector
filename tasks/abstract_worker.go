package tasks

import "sync"

type ConnectionWorker interface {
	Start(group *sync.WaitGroup)
}

type BackGroundWorker interface {
	Start()
}
