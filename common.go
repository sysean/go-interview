package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下
// 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
func forGoTest01() {
	fmt.Println("forGoTest01 start")
	a2BChan := make(chan struct{})
	b2AChan := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		i := 1
		for i = 1; i <= 28; {
			select {
			case <-b2AChan:
				fmt.Printf("%d", i)
				i++
				fmt.Printf("%d", i)
				i++
				a2BChan <- struct{}{}
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 65; i <= 90; {
			select {
			case <-a2BChan:
				fmt.Printf("%c", i)
				i++
				fmt.Printf("%c", i)
				i++
				b2AChan <- struct{}{}
			}
		}
	}()

	b2AChan <- struct{}{}
	wg.Wait()
	fmt.Println("")
	fmt.Println("forGoTest01 end")
	fmt.Println("")
}

// 会panic，wg计数达到了负数
func forGoTest02() {
	fmt.Println("forGoTest02 start")
	wg := sync.WaitGroup{}
	wg1 := wg // 编译期，WaitGroup 不能拷贝
	wg1.Done()
	c := make(chan struct{})
	wg.Add(1)

	go func() {
		defer wg.Done()
		<-c
		fmt.Println("go func done")
	}()

	time.Sleep(3 * time.Second)
	wg.Add(-1)
	time.Sleep(3 * time.Second)
	fmt.Println("wg done")
	fmt.Println("forGoTest02 end")
	fmt.Println("")
}

func forGoTest03() {
	fmt.Println("forGoTest03 start")
	const timeout = 5 * time.Second

	f1 := func() string {
		time.Sleep(10 * time.Second)
		return "f1"
	}

	f2 := func() string {
		time.Sleep(5 * time.Second)
		return "f2"
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resultChan := make(chan string)

	start := time.Now()

	go func() {
		select {
		case <-ctx.Done(): // 这里是为了节约开销，如果调度器已经调用了 cancel，那么就不需要再调用 f1 了
			fmt.Println("f1 done")
		default:
			resultChan <- f1()
		}
	}()

	go func() {
		select {
		case <-ctx.Done(): // 这里是为了节约开销，如果调度器已经调用了 cancel，那么就不需要再调用 f2 了
			fmt.Println("f2 done")
		default:
			resultChan <- f2()
		}
	}()

	for i := 0; i < 2; i++ {
		select {
		case <-ctx.Done(): // 这里是为了监测是否超时
			fmt.Println("timeout")
			break
		case result := <-resultChan:
			fmt.Println(result)
		}
	}

	fmt.Println("over, use time: ", time.Since(start))
	fmt.Println("forGoTest03 end")
	fmt.Println("")
}
