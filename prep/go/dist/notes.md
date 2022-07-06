# Example 1
Usually, the final value of the loop variable (10) is printed by all of the goroutines (this is dependent on when the goroutines launch relative to the progress of the loop. I sometimes see intermediate values when running this).

The go compiler outlines (opposite of _inlines_) the closure and passes the values that it references from its scope to the outlined function by reference.

```go
func main() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("launched goroutine %d\n", i)
		}()
	}
	// Wait for goroutines to finish
	time.Sleep(time.Second)
}
```

is translated to

```go
func func1(i *int) {
	fmt.Printf("launched goroutine %d\n", *i)
}

func main() {
	for i := 0; i < 10; i++ {
		go func1(&i)
	}
	// Wait for goroutines to finish
	time.Sleep(time.Second)
}
```

Note that by the go spec, the loop variable is re-used, so all go routines will receive the same pointer.

To fix this, we can explicitly pass the loop variable as an argument to the anonymous function.


```go
func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("launched goroutine %d\n", i)
		}(i)
	}
	// Wait for goroutines to finish
	time.Sleep(time.Second)
}
```

[This post](https://eli.thegreenplace.net/2019/go-compiler-internals-adding-a-new-statement-to-go-part-1/) was very useful for understanding this problem.

# Example 2

```go
package main

import (
	"fmt"
)

const numTasks = 3

func main() {
	var done chan struct{}
	for i := 0; i < numTasks; i++ {
		go func() {
			fmt.Println("running task...")

			// Signal that task is done
			done <- struct{}{}
		}()
	}

	// Wait for tasks to complete
	for i := 0; i < numTasks; i++ {
		<-done
	}
	fmt.Printf("all %d tasks done!\n", numTasks)
}
```

This prints out:

```
running task...
running task...
running task...
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive (nil chan)]:
main.main()
        /Users/panashe/code/panashe-csi/prep/go/dist/example-2.go:22 +0x6a

goroutine 6 [chan send (nil chan)]:
main.main.func1()
        /Users/panashe/code/panashe-csi/prep/go/dist/example-2.go:16 +0x6a
created by main.main
        /Users/panashe/code/panashe-csi/prep/go/dist/example-2.go:12 +0x49

goroutine 7 [chan send (nil chan)]:
main.main.func1()
        /Users/panashe/code/panashe-csi/prep/go/dist/example-2.go:16 +0x6a
created by main.main
        /Users/panashe/code/panashe-csi/prep/go/dist/example-2.go:12 +0x49

goroutine 8 [chan send (nil chan)]:
main.main.func1()
        /Users/panashe/code/panashe-csi/prep/go/dist/example-2.go:16 +0x6a
created by main.main
        /Users/panashe/code/panashe-csi/prep/go/dist/example-2.go:12 +0x49
exit status 2
```

This deadlocks because a receive from a `nil` channel blocks forever. This
channel was not initialized, which is where the channel is given a buffer size,
and a `nil` channel is zero-sized. Both send and receive will block on a
zero-sized channel.
See [channel axiom 2](https://dave.cheney.net/2014/03/19/channel-axioms).

To fix:

```go
func main() {
	var done chan struct{} = make(chan struct{}, 3)
	for i := 0; i < numTasks; i++ {
		go func() {
			fmt.Println("running task...")

			// Signal that task is done
			done <- struct{}{}
		}()
	}

	// Wait for tasks to complete
	for i := 0; i < numTasks; i++ {
		<-done
	}
	fmt.Printf("all %d tasks done!\n", numTasks)
}
```


# Example 3

```go
package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var urls = []string{
		"https://bradfieldcs.com/courses/architecture/",
		"https://bradfieldcs.com/courses/networking/",
		"https://bradfieldcs.com/courses/databases/",
	}
	var wg sync.WaitGroup
	for i := range urls {
		go func(i int) {
			wg.Add(1)
			// Decrement the counter when the goroutine completes.
			defer wg.Done()

			_, err := http.Get(urls[i])
			if err != nil {
				panic(err)
			}

			fmt.Println("Successfully fetched", urls[i])
		}(i)
	}

	// Wait for all url fetches
	wg.Wait()
	fmt.Println("all url fetches done!")
}
```

prints only:

```
all url fetches done!
```

from the documentation of [`Waitgroup.Add`](https://pkg.go.dev/sync@go1.18.3#WaitGroup.Add):

> Note that calls with a positive delta that occur when the counter is zero must happen before a Wait. Calls with a negative delta, or calls with a positive delta that start when the counter is greater than zero, may happen at any time. Typically this means the calls to Add should execute before the statement creating the goroutine or other event to be waited for.

So the code above presents a subtle race condition. The first goroutine may not launch, and thus call `wg.Add(1)` before the call to `wg.Wait()`. The fix is to move the `Add` call before the goroutine.

```go
package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var urls = []string{
		"https://bradfieldcs.com/courses/architecture/",
		"https://bradfieldcs.com/courses/networking/",
		"https://bradfieldcs.com/courses/databases/",
	}
	var wg sync.WaitGroup
	for i := range urls {
		wg.Add(1)
		go func(i int) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()

			_, err := http.Get(urls[i])
			if err != nil {
				panic(err)
			}

			fmt.Println("Successfully fetched", urls[i])
		}(i)
	}

	// Wait for all url fetches
	wg.Wait()
	fmt.Println("all url fetches done!")
}
```

# Example 4
```go
package main

import (
	"fmt"
)

func main() {
	done := make(chan struct{}, 1)
	go func() {
		fmt.Println("performing initialization...")
		<-done
	}()

	done <- struct{}{}
	fmt.Println("initialization done, continuing with rest of program")
}
```

Here, the output is:

```
initialization done, continuing with rest of program
```

The issue here is that `main` returns before the goroutine is executed. Note that receive from a channel blocks until a value becomes available on that channel. Since we're using this channel purely to communicate status between main and the goroutine, we should receive in main and put in the goroutine instead.

```go
package main

import (
	"fmt"
)

func main() {
	done := make(chan struct{}, 1)
	go func() {
		fmt.Println("performing initialization...")
		done <- struct{}{}
	}()


	<-done
	fmt.Println("initialization done, continuing with rest of program")
}
```


# Example 5 (unsolved)
```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var responses = []string{
	"200 OK",
	"402 Payment Required",
	"418 I'm a teapot",
}

func randomDelay(maxMillis int) time.Duration {
	return time.Duration(rand.Intn(maxMillis)) * time.Millisecond
}

func query(endpoint string) string {
	// Simulate querying the given endpoint
	delay := randomDelay(100)
	time.Sleep(delay)

	i := rand.Intn(len(responses))
	return responses[i]
}

// Query each of the mirrors in parallel and return the first
// response (this approach increases the amount of traffic but
// significantly improves "tail latency")
func parallelQuery(endpoints []string) string {
	results := make(chan string)
	for i := range endpoints {
		go func(i int) {
			results <- query(endpoints[i])
		}(i)
	}
	return <-results
}

func main() {
	var endpoints = []string{
		"https://fakeurl.com/endpoint",
		"https://mirror1.com/endpoint",
		"https://mirror2.com/endpoint",
	}

	// Simulate long-running server process that makes continuous queries
	for {
		fmt.Println(parallelQuery(endpoints))
		delay := randomDelay(100)
		time.Sleep(delay)
	}
}
```

It's not immediately evident what the issue is here. Let's modify this to make the problem more apparent:

```go
func parallelQuery(endpoints []string) []string {
	results := make(chan []string)
	for i := range endpoints {
		go func(i int) {
			results <- []string{endpoints[i], query(endpoints[i])}
		}(i)
	}
	return <-results
}

func main() {
	var endpoints = []string{
		"https://fakeurl.com/endpoint",
		"https://mirror1.com/endpoint",
		"https://mirror2.com/endpoint",
	}

	// Simulate long-running server process that makes continuous queries
	for i := 0; i < 5; i++ {
		fmt.Println(parallelQuery(endpoints))
		delay := randomDelay(100)
		time.Sleep(delay)
	}
}
```

I find that mirror1 is usually the chosen server (stable across runs).
```
[https://mirror1.com/endpoint 418 I'm a teapot]
[https://mirror1.com/endpoint 418 I'm a teapot]
[https://mirror1.com/endpoint 200 OK]
[https://mirror1.com/endpoint 200 OK]
[https://mirror1.com/endpoint 402 Payment Required]
```


# Example 6
```go
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

func (c *coordinator) logState() {
	c.lock.RLock()
	defer c.lock.RUnlock()

	fmt.Printf("leader = %q\n", c.leader)
}

func (c *coordinator) setLeader(leader string, shouldLog bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.leader = leader

	if shouldLog {
		c.logState()
	}
}

func main() {
	c := newCoordinator("us-east")
	c.logState()
	c.setLeader("us-west", true)
}
```

when run, produces:

```
leader = "us-east"
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_SemacquireMutex(0xc000078ec0, 0x13, 0xc000078ec0)
        /nix/store/1168hmr97r5lk8d7y0hinp3vf42sq4r2-go-1.17.7/share/go/src/runtime/sema.go:71 +0x25
sync.(*RWMutex).RLock(...)
        /nix/store/1168hmr97r5lk8d7y0hinp3vf42sq4r2-go-1.17.7/share/go/src/sync/rwmutex.go:63
main.(*coordinator).logState(0xc00007c180)
        /Users/panashe/code/panashe-csi/prep/go/dist/example-6.go:21 +0x4e
main.(*coordinator).setLeader(0xc00007c180, {0x10a3f6a, 0x7}, 0x1)
        /Users/panashe/code/panashe-csi/prep/go/dist/example-6.go:34 +0xa9
main.main()
        /Users/panashe/code/panashe-csi/prep/go/dist/example-6.go:41 +0x65
exit status 2
```

The `logState` method is attempting to acquire the read lock while the write lock is already held. Three possible solutions are to:
- rely on the caller to correctly lock before calling `logState`
- pass `logState` a boolean to tell it whether to attempt to acquire the lock
- split our `logState` method into two variants, one that does acquire locks, and one that doesn't. With this design, the one that acquires locks could be public, while the other is kept private.

An example of the third solution:

```go
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

func (c *coordinator) logState() {
	c.lock.RLock()
	defer c.lock.RUnlock()

	fmt.Printf("leader = %q\n", c.leader)
}

func (c *coordinator) setLeader(leader string, shouldLog bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.leader = leader

	if shouldLog {
		c.logState()
	}
}

func main() {
	c := newCoordinator("us-east")
	c.logState()
	c.setLeader("us-west", true)
}
```

# Example 7
```go
package main

import (
	"fmt"
	"sync"
)

const (
	numGoroutines = 100
	numIncrements = 100
)

type counter struct {
	count int
}

func safeIncrement(lock sync.Mutex, c *counter) {
	lock.Lock()
	defer lock.Unlock()

	c.count += 1
}

func main() {
	var globalLock sync.Mutex
	c := &counter{
		count: 0,
	}

	var wg sync.WaitGroup
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < numIncrements; j++ {
				safeIncrement(globalLock, c)
			}
		}()
	}

	wg.Wait()
	fmt.Println(c.count)
}
```

The expected value of the print out is \(100 * 100 = 10000\) in the first run on my machine, this printed 5393.

The issue here is that `safeIncrement` receives a copy of `lock`. In Go, primitive values and structs that only contain primitives are copied when passed as function arguments, and `sync.Mutex` fits this description. The solution is to pass a reference to `safeIncrement`. Note that I'm relying on go's implicit pointer dereference.

```go
package main

import (
	"fmt"
	"sync"
)

const (
	numGoroutines = 100
	numIncrements = 100
)

type counter struct {
	count int
}

func safeIncrement(lock *sync.Mutex, c *counter) {
	lock.Lock()
	defer lock.Unlock()

	c.count += 1
}

func main() {
	var globalLock sync.Mutex
	c := &counter{
		count: 0,
	}

	var wg sync.WaitGroup
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < numIncrements; j++ {
				safeIncrement(&globalLock, c)
			}
		}()
	}

	wg.Wait()
	fmt.Println(c.count)
}
```

# Example 8
# Example 9
