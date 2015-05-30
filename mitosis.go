package main

import (
	"fmt"
	"mitosis/container"
)

func main(){
	fmt.Println("hello world")
	c:=&container.Container{}
	c.Start()
}
