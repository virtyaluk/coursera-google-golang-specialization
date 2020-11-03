package main

import (
	"fmt"
	"time"
)

func race(str string) {
	fmt.Println(str)
}

func main() {
	for i := 0; i < 5; i++ {
		go race("foo")
		go race("bar")

		time.Sleep(2 * time.Millisecond)
	}
}

/*

The program consist of a simple loop that executes 2 goroutines
each iteration and prints strings "foo" and "bar".
The print messages appear in a wrong order thus introducing so-called race condition.

The output below demonstrates that the printed messages
appear in a different order each time:

~/De/projects/courses/coursera-google-golang/golang-concurrency/week-2 master ?6 ❯ go build main.go                                                                                 06:48:59 PM
~/De/projects/courses/coursera-google-golang/golang-concurrency/week-2 master ?6 ❯ ./main                                                                                           06:58:56 PM
foo
bar
bar
foo
foo
bar
bar
foo
bar
foo
~/De/projects/courses/coursera-google-golang/golang-concurrency/week-2 master ?6 ❯ ./main                                                                                           06:58:59 PM
foo
bar
bar
foo
bar
foo
bar
foo
bar
foo
~/De/projects/courses/coursera-google-golang/golang-concurrency/week-2 master ?6 ❯ ./main                                                                                           06:59:03 PM
bar
foo
foo
bar
foo
bar
bar
foo
bar
foo

*/
