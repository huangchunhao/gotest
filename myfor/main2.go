package main

import "fmt"

func main() {
	count := 1
	str := "abcde"
	str2 := "12345"
	for _, data := range str {
		//fmt.Println(data)
		for _, data2 := range str2 {
			fmt.Printf("%d %c%c\n", count, data, data2)
			count += 1
		}
	}

	//百钱百鸡
	//一只公鸡5钱  一只母鸡3钱   三只小鸡1钱
	count1:=0
	for cock := 0; cock <= 20; cock++ {
		for hen := 0; hen <= 33; hen++ {
			for chicken := 0; chicken <= 100; chicken += 3 {
				count1++
				if cock+hen+chicken == 100 && cock*5+hen*3+chicken/3 == 100 {
					fmt.Printf("%d %d %d\n", cock, hen, chicken)
				}
			}
		}
	}
	fmt.Println("第一种循环次数")
	fmt.Println(count1)


	count2:=0
	for cock := 0; cock <= 20; cock++ {
		for hen := 0; hen <= 33; hen++ {
			chicken := 100 - hen - cock
			count2++
			if chicken % 3 == 0 && cock*5+hen*3+chicken/3 == 100 {
				fmt.Printf("%d %d %d\n", cock, hen, chicken)
			}

		}
	}
	fmt.Println("第二种循环次数")
	fmt.Println(count2)


	//0 25 75
	//4 18 78
	//8 11 81
	//12 4 84
	//第一种循环次数
	//24276
	//0 25 75
	//4 18 78
	//8 11 81
	//12 4 84
	//第二种循环次数
	//714
}
