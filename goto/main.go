package main

import (
	"fmt"
	"time"
)

func main() {
	i:=0
	for{
		i++
		time.Sleep(time.Second)
		fmt.Println(i)
		if i==3{
			goto FLAG//当为3的时候直接执行FLAG后面的语句
		}
		if i==5 {
			fmt.Println("break跳出,结束")
			break
		}
		if i==2 {
			fmt.Println("continue跳出本次循环")
			continue//  只有在for循环中使用     continue后面的语句不再执行
		}
		fmt.Println("---------")
		FLAG:
		fmt.Println("+++++++++")
	}
	
}
