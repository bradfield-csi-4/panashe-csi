package main

import (
	"fmt"
	"sync"
)

type coordinator struct {
	lock   sync.RWMutex
	leader string
}

func newCoordinator(leader string) *coordinator {
	return &coordinator{
		lock:   sync.RWMutex{},
		leader: leader,
	}
}

func (c *coordinator) LogState() {
	c.lock.RLock()
	defer c.lock.RUnlock()
	c.unsafeLogState()
}

// unsafeLogState should only be called once the coordinator lock is held, to prevent reporting stale data.
func (c *coordinator) unsafeLogState() {
	fmt.Printf("leader = %q\n", c.leader)
}

func (c *coordinator) setLeader(leader string, shouldLog bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.leader = leader

	if shouldLog {
		// Don't attempt to acquire the lock, write lock already held.
		c.unsafeLogState()
	}
}

func main() {
	c := newCoordinator("us-east")
	c.LogState()
	c.setLeader("us-west", true)
}
