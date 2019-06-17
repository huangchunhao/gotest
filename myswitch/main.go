package main

import "fmt"

func main() {

	//var score int =20
	score := 20
	//不需要break
	switch score {

	case 90:
		fmt.Println("A")
		fallthrough //强制再执行下一个case
	case 80, 70:
		fmt.Println("B")
	case 60, 50, 40:
		fmt.Println("C")
	default:
		fmt.Println("default")
	}
	//switch 方式1
	switch s1 := 90; s1 {
	case 90:
		fmt.Println("A")
		fallthrough //强制再执行下一个case
	case 80, 70:
		fmt.Println("B")
	case 60, 50, 40:
		fmt.Println("C")
	default:
		fmt.Println("default")
	}
	//switch 方式2
	switch s3 := 90; {
	case s3 >= 90:
		fmt.Println("A")
	case s3 >= 80:
		fmt.Println("B")
	case s3 >= 70:
		fmt.Println("C")
	default:
		fmt.Println("default")
	}
	//switch 方式3
	var s2 int = 80
	switch {
	case s2 >= 90:
		fmt.Println("A")
	case s2 >= 80:
		fmt.Println("B")
	case s2 >= 70:
		fmt.Println("C")
	default:
		fmt.Println("default")
	}
}
