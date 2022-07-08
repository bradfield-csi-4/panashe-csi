package main

import (
	"fmt"
	"sync"
	"time"
)

type dbService struct {
	lock       *sync.RWMutex
	connection string
}

func newDbService(connection string) *dbService {
	return &dbService{
		lock:       &sync.RWMutex{},
		connection: connection,
	}
}


// unsafeLogState should only be called while holding a read lock to avoid stale reads.
func (d *dbService) unsafeLogState() {
	fmt.Printf("connection %q is healthy\n", d.connection)
}

func (d *dbService) takeSnapshot() {
	d.lock.RLock()
	defer d.lock.RUnlock()

	fmt.Printf("Taking snapshot over connection %q\n", d.connection)

	// Simulate slow operation
	time.Sleep(time.Second)

	d.unsafeLogState()
}

func (d *dbService) updateConnection(connection string) {
	d.lock.Lock()
	defer d.lock.Unlock()

	d.connection = connection
}

func main() {
	d := newDbService("127.0.0.1:3001")

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		d.takeSnapshot()
	}()

	// Simulate other DB accesses
	time.Sleep(200 * time.Millisecond)

	wg.Add(1)
	go func() {
		defer wg.Done()

		d.updateConnection("127.0.0.1:8080")
	}()

	wg.Wait()
}
