package main

import (
	"fmt"
	"sync"
)
type Info struct {
    mu sync.Mutex
    // ... other fields, e.g.: Str string
    Str string
}
//可上锁的共享缓冲器
type SyncedBuffer struct {
    lock    sync.Mutex
    buffer  bytes.Buffer
}
func main() {
	test := Info {Str: "aa"}
	fmt.Println(test.Str)
	test.Str = "aaa"
	fmt.Println(test.Str)
	Update(&test)
	fmt.Println(test.Str)
}
func Update(info *Info) {
    info.mu.Lock()
    // critical section:
    info.Str = "aaaa"// new value
    // end critical section
    info.mu.Unlock()
}