package main

import (
	"fmt"
)

func addition(x int, y int) int {
	return x + y;
}

func sayHello(s string) string {
	var b bool  //默认布尔为false， go中不能用0，1代替布尔类型
	fmt.Println(b)
	b = true
	fmt.Println(b)
	return "Hello " + s
}

func main()  {
   fmt.Println(sayHello("Liuwei"))
}