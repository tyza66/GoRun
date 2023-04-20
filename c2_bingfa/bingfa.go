package main

/**
 * Author: tyza66
 * Github: https://github.com/tyza66
 **/
import (
	"fmt"
	"time"
)

func main() {
	//go程 携程
	//Go 程在相同的地址空间中运行，因此在访问共享的内存时必须进行同步
	//go say("Hello")
	//say("giao")

	//信道
	ch := make(chan int)
	//如果携程中的值一直没给过来 那么另外的携程就会阻塞等待 如果阻塞等待无法被解开就会进入死锁状态
	i1 := []int{1, 2, 3, 4, 5, 6}
	//i2 := []int{4,5,6}
	go sum(i1[:len(i1)/2], ch)
	go sum(i1[len(i1)/2:], ch)
	x, y := <-ch, <-ch
	fmt.Println(x, y)
	//带缓冲的信道 生成信道的时候可以说明缓冲带的长度 如果信道的缓冲带满了 那么输入信道的携程才会阻塞 当缓存区为空时接收信道的携程才会阻塞
	ch1 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2
	//ch1<-3 多这个会造成阻塞 因为信道满了
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
	//fmt.Println(<-ch1) 多这个也会造成阻塞 因为信道里面空了
	//信道的关闭和状态检测
	ch1 <- 4
	ch1 <- 3
	a, ok := <-ch1
	fmt.Println(a, ok)
	b, ok1 := <-ch1
	fmt.Println(b, ok1)

	close(ch)
	close(ch1)

	fmt.Println()
	c3 := make(chan int, 10)
	//使用range循环能自动执行到信道被关闭
	go fibonacci(cap(c3), c3)
	for i := range c3 {
		fmt.Println(i)
	}

	//select 语句使一个 Go 程可以等待多个通信操作。
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		//这个斐波那契执行完之后也让那个停下来
		quit <- 0
	}()
	fibonacci2(c, quit)
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	//这个时候把sum的值传到了信道中
	c <- sum
}

//这个是一个斐波那契函数
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

//这个斐波那契实现的是一个go程等待多个信道可用 并且select会阻塞到某个分支可以继续执行为止
//当多个分支都准备好了的时候会随机选择一个执行
//当 select 中的其它分支都没有准备好时，default 分支就会执行
func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
