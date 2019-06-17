package main

import "fmt"

func main(){
	var define_s ,define_s1 string
	define_s,define_s1="45662","546565"
	define_s2 := "565454"
	var define_s3 string ="87788"


	var f_d float64=2
	var f_d_1 float64=f_d*f_d

	fmt.Println(define_s,define_s1,define_s2,define_s3)
	fmt.Println(f_d_1)
	fmt.Println("1233")

	//多重变量
	a,b,c :=10,20,30
	fmt.Println(a,b,c)
	a,b = b,a
	fmt.Println(a,b,c)
	// _匿名变量
	temp,_:=re()
	fmt.Println(temp)

	//var (
	//	e =1
	//	f ="555"
	//	)
	e,f:=1,"555"
	fmt.Println(e,f)


	const  j int=5689
	const (
		h string="5986"
		i float64=3.14
	)
	fmt.Println(h,i)
	const (
		k ="5986"
		n =3.14
	)
	fmt.Println(k,n)

	//iota常量自动生成器   每个一行  自动化加1
	//iota给常量赋值使用
	//iota遇到const  重置为0
	//同一个小括号中，可以使用一个iota
	const(
		ai= iota
		bi=iota
		ci=iota
	)
	fmt.Println(ai,bi,ci)

	const di=iota
	fmt.Println(ai,bi,ci,di)

	const(
		a1= iota
		b1
		c1
	)
	fmt.Println(a1,b1,c1)
//这里不能省略iota
	const(
		ai1,ai2= iota,iota
		bi1=iota
		ci1=iota
	)
	fmt.Println(ai1,ai2,bi1,ci1)

}

func re() (a int,b string){
	return 1,"2"
}
