package main

import (
	"fmt"
	"strconv"
)

func Add(ch chan int, i int) {
	fmt.Println("<<<<<<add" + strconv.Itoa(i))
	ch <- i
}

func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		s:="main" + strconv.Itoa(i)
		fmt.Println(s)
		chs[i] = make(chan int)
		go Add(chs[i],i)
	}

	for _, ch := range (chs) {
		i:=<-ch
		sum := 1+i
		fmt.Println("<<<<<<sum" + strconv.Itoa(sum))
		fmt.Println("chu>>>>>>>" + strconv.Itoa(i))
	}
}


