package main

import (
	"fmt"
	"sort"
)

func main() {

	//变量声明
	var myMap map[string] string
	//newmymap :=make(map[string] string)

	//创建
	//myMap=map[string] string{
	//	"1":"1",
	//	"2":"2"}

	//myMap=make(map[string] string)
	myMap=make(map[string] string,100)
	//赋值
	myMap["1"]="1"
	myMap["2"]="2"
	myMap["3"]="3"
	myMap["del"]="del"

	for k, v := range myMap {
		fmt.Println(k,v)
	}
	//删除
	delete(myMap,"del")
	for k, v := range myMap {
		fmt.Println(k,v)
	}
	//查找
	value,ok:=myMap["3"]
	if ok {
		fmt.Println(value)
	}
	//排序-通过数组
	var sort_array =make([]string,len(myMap))
	for k, _ := range myMap {
		sort_array=append(sort_array, k)
	}
	sort.Strings(sort_array)
	for _, v := range sort_array {
		fmt.Println(v,myMap[v])
	}

}
