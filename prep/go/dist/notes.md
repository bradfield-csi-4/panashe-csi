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

# Example 4
# Example 5
# Example 6
# Example 7
# Example 8
# Example 9
