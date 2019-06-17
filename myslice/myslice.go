package main

import "fmt"

func main() {
	//var myArray [10]int=[10]int{1,2,3,4,5,6,7,8,9,10}
	//或者
	//myArray := []int{1,2,3,4,5,6,7,8,9,10}
	//或者
	myArray := make([]int, 5, 10)
	fmt.Println(cap(myArray))
	fmt.Println(len(myArray))
	myArray = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice1_5 := myArray[:5]
	fmt.Println(cap(slice1_5))
	fmt.Println(len(slice1_5))

	//遍历
	for _, v := range slice1_5 {
		fmt.Println(v)
	}

	slice1_1 := myArray[:1]
	for _, v := range slice1_1 {
		fmt.Println(v)
	}

	slice5_10 := myArray[5:]
	for _, v := range slice5_10 {
		fmt.Println(v)
	}

	//append后尾添加
	newmyArray := append(myArray, 11, 12, 13)
	newmyArray2 := append(myArray, slice5_10...)
	for _, v := range newmyArray {
		fmt.Println(v)
	}
	for _, v := range newmyArray2 {
		fmt.Println(v)
	}

	//copy复制
	myArraynew := make([]int, 5, 10)
	copy(myArraynew, myArray)
	for _, v := range newmyArray2 {
		fmt.Println(v)
	}
}
