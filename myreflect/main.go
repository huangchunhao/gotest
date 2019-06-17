package main

import (
	"fmt"
	"reflect"
)

type aaa int

type aaaa aaa

type Bbb struct {
	Ccc string
	Ddd int `tag ddd`
}

func add(a, b int) int {
	return a + b
}

func main() {
	var aaa aaaa
	type0fA := reflect.TypeOf(aaa)
	fmt.Println(type0fA.Name(), type0fA.Kind())
/////////////////////////////////////
//结构体相关处理
	bbb := new(Bbb)
	bbb.Ccc = "123"
	bbb.Ddd = 123

	type0fStruct := reflect.TypeOf(*bbb)
	for i := 0; i < type0fStruct.NumField(); i++ {
		// 获取每个成员的结构体字段类型
		fieldType := type0fStruct.Field(i)
		// 输出成员名和tag
		fmt.Printf("name: %v  tag: '%v'\n", fieldType.Name, fieldType.Tag)
	}

	// 值包装结构体
	d := reflect.ValueOf(*bbb)

	// 获取字段数量
	fmt.Println("NumField", d.NumField())
	// 获取索引为1的字段(int字段)
	floatField0 := d.Field(0)
	floatField1 := d.Field(1)
	// 输出字段类型
	fmt.Println("Field:", floatField0.Type(),"Value:",floatField0.Interface().(string))
	fmt.Println("Field:", floatField1.Type(),"Value:",floatField1.Interface().(int))
	// 根据名字查找字段
	fmt.Println("FieldByName(\"Ccc\").Type", d.FieldByName("Ccc").Type())
	// 根据索引查找值中, next字段的int字段的值
	//fmt.Println("FieldByIndex([]int{1, 0}).Type()", d.FieldByIndex([]int{1, 0}).Type())

	/////////////////////////////
	// 声明整型变量a并赋初值
	var a int = 1024
	// 获取变量a的反射值对象
	valueOfA := reflect.ValueOf(a)
	// 获取interface{}类型的值, 通过类型断言转换
	var getA int = valueOfA.Interface().(int)
	// 获取64位的值, 强制类型转换为int类型
	var getA2 int = int(valueOfA.Int())
	fmt.Println(getA, getA2)

	/////////////////////////
	//反射调用函数
	// 将函数包装为反射值对象
	funcValue := reflect.ValueOf(add)
	// 构造函数参数, 传入两个整型值
	paramList := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}
	// 反射调用函数
	retList := funcValue.Call(paramList)
	// 获取第一个返回值, 取整数值
	fmt.Println(retList[0].Int())
}
