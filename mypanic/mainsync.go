package main

import (
	"fmt"
	"sync"
)

func base2(init string,num int) string{
	fmt.Println("begin",num)
	ss:=init[0:(len(init)-num)]
	fmt.Println("end",num)
	return ss
}

func yu(init string,num int) string{
	defer func() {
		//recover()：返回程序为什么挂了
		if err := recover(); err != nil {
			fmt.Println("-------")
			fmt.Println(err)
			fmt.Println("-------")
		}
	}()

	 sss:=base2(init,num);
	 //fmt.Println("只会影响调用base方法的逻辑块")

	return sss;
}

func run(i int,wg *sync.WaitGroup,init *string){
	if i==5{
		//fmt.Println(i)
		yu(*init,19)
	}else{
		//fmt.Println(i)
		yu(*init,7)
	}
	wg.Done()
}



func main(){
	//defer fmt.Println("over")

	wg:=sync.WaitGroup{}
	wg.Add(10)
	init:="1234567890"


	for i:=1;i<11;i++{
		go run(i,&wg,&init)
		//如果yu方法中没有恢复代码，这里单个线程如果报错，整个main将停止执行。有了之后就可以不影响其他线程执行，可以保证main执行完整。
	}
	wg.Wait()





}
