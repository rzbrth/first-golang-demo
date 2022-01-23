package main

import (
	"fmt"
	"sync"
	"time"
)

func greetuserFromOutside(){
	fmt.Print("Hello from other side within same package..........\n")
}

func greetuserFromOutsideWithSleep(wg *sync.WaitGroup){
	time.Sleep(30* time.Second)
	fmt.Print("Hello from other side within same package and sleeping ..........\n")

	wg.Done()
}