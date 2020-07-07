package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	i := 10000
	for ; i < 12000; i++ {
		sprint := fmt.Sprintf("\"%s\",\"%s\",%s", strconv.Itoa(i), "25d55ad283aa400af464c76d713c07ad", strconv.Itoa(i))
		fmt.Println(sprint)
	}
	fmt.Println(time.Now().Unix())
	defer_call()

}



func defer_call() {
	defer func() {
		fmt.Println("打印前")
	}()
	panic("触发异常")
	defer func() {
		fmt.Println("打印中")
	}()
	defer func() {
		fmt.Println("打印后")
	}()

}