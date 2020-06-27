package main

import "fmt"

func base(init string,num int) string{
	return init[0:(len(init)-num)]
}

func call(init string,num int) string{
	defer func() {
		//recover()：返回程序为什么挂了
		if err := recover(); err != nil {
			fmt.Println("-------")
			fmt.Println(err)
			fmt.Println("-------")
		}
	}()

	 sss:=base(init,num);
	  fmt.Println("只会影响调用base方法的逻辑块")

	return sss;
}

func main(){
	//defer fmt.Println("over")

	init:="1234444i"
	fmt.Println(call(init,10))
	fmt.Println("不影响调用call方法的逻辑块，也就是说本逻辑块被恢复了")


}
