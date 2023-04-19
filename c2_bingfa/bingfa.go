package main

import (
	"fmt"
	"time"
)

func main() {
	//go程
	go say("Hello")
	say("giao")
}

func say(s string){
	for i := 0;i<5;i++{
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}
}