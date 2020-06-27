package main

import "fmt"

//系统抛异常
func test0111() {
	a := [5]int{0, 1, 2, 3, 4}
	a[1] = 123
	//a[10] =456
	index := 10
	//index out of range
	a[index] = 456
}

//自己抛
func test222() {
	getCircleArea(-5)
}

func getCircleArea(radius float32) (area float32) {
	if radius < 0 {
		//自己抛异常
		panic("您的智商已下线，半径不能为负数")
	}
	return 3.14 * radius * radius
}

//结合defer
func test0333() {
	//延迟执行匿名函数
	//(1)函数正常执行结束了
	//(2)报错中断执行
	defer func() {
		//recover()：返回程序为什么挂了
		if err := recover(); err != nil {
			fmt.Println("-------")
			fmt.Println(err)
			fmt.Println("-------")
		}
	}()
	getCircleArea(-5)
	fmt.Println("这里有没有执行")
}

func test444() {
	test0333()
	fmt.Println("GAME OVER")
}

func main() {
	//test0111()
	test222()
	//test0333()
	//test444()
}