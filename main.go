package main

import (
	"sync"

	"github.com/senaduka/cetvorke/cetgemini"
)

func main() {
	waitgroup := &sync.WaitGroup{}
	waitgroup.Add(1)
	go cetgemini.StartServer(waitgroup)
	waitgroup.Wait()
}
