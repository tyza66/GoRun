package main

import (
	"fmt"
	"time"
)

func main() {
	//go程 携程
	//Go 程在相同的地址空间中运行，因此在访问共享的内存时必须进行同步
	go say("Hello")
	say("giao")
}

func say(s string){
	for i := 0;i<5;i++{
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}
}