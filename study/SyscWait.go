package study

import (
	"fmt"
	"sync"
	"time"
)

func TestName() {
	// 定义WaitGroup
	var wg sync.WaitGroup

	// 计数器+1
	wg.Add(1)
	go foo(&wg, "a")

	time.Sleep(time.Second * 3)

	// 计数器+1
	wg.Add(1)
	go foo(&wg, "b")

	fmt.Println("before wait")

	// 等待WaitGroup计数器归零
	wg.Wait()

	fmt.Println("after wait")
}

func foo(wg *sync.WaitGroup, name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(name, i)
		time.Sleep(time.Second)
	}
	// 计数器-1
	wg.Done()
}
