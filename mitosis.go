package main

import (
	"fmt"
	"mitosis/container"
	"time"
)

func main(){
	fmt.Println("hello world")
	c:=&container.Container{Hostname:"testContainer"}
	c.Start()
	time.Sleep(10)
	fmt.Println("pid",c.Pid)
}
