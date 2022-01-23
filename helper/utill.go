package helper

import (
	"fmt"
	"sync"
	"time"
)

//Global or package level variable
var ConferenceName = "Go Conference" //variable export worthy

//Making first letter capital make the function export worthy i.e. we can use it outside the package
func GreetuserFromOutside(wg *sync.WaitGroup){
	time.Sleep(10 * time.Second)
	fmt.Print("Hello from other side from other package..........\n")
	wg.Done()
}