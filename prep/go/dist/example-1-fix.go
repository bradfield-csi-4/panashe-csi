package main

import (
	"fmt"
	"time"
)

/*
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
*/
func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("launched goroutine %d\n", i)
		}(i)
	}
	// Wait for goroutines to finish
	time.Sleep(time.Second)
}
