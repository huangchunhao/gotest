package main

import "fmt"

func main() {
	sum :=0

	for i:=1;i<11;i++{
		sum=sum+i
	}
	fmt.Println(sum)

	str:="abc"
	for i:=0;i<len(str);i++{
		fmt.Println(str[i])
		fmt.Printf("str[%d]=%c\n",i,str[i])
	}

	for i,data:=range str{
		fmt.Println(i,data)
		fmt.Printf("str[%d]=%c\n",i,data)
	}

	for _,data:=range str{
		fmt.Println(data)
		fmt.Printf("%c\n",data)
	}
}
