package main

import "fmt"

type Point struct {
	x, y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func Prp(p *Point) {
	fmt.Println(p)
}

func main() {
	p := Point{1, 2}
	fmt.Printf("%T\n",p)
	//没有调用函数调用
	fmt.Println(p)

	//赋值后打印
	Prp(&p)
	//直接打印
	Prp(&Point{2, 3})

	//new返回的是指针类型
	pp:=new(Point)
	*pp=Point{4, 5}
	fmt.Println(pp)
	fmt.Printf("%T\n",pp)
	Prp(pp)


	//匿名成员-可以直接访问，不需要写出完整路径（注意type的定义方式 以及嵌入方式）
	var w Wheel
	w.x=8
	w.y=8
	w.Radius=5
	w.Spokes=20
	fmt.Println(w)
	fmt.Println(w.Spokes)
	fmt.Println(w.Circle)
	fmt.Println(w.Point)
	Prp(&w.Point)
}
