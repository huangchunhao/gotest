package main

import "fmt"

//func MustCompile(str string) *Regexp {
//	regexp, error := Compile(str)
//	if error != nil {
//		panic(`regexp: Compile(` + quote(str) + `): ` + error.Error())
//	}
//	return regexp
//}

func main() {
	defer fmt.Println("宕机后要做的事情1")
	defer fmt.Println("宕机后要做的事情2")
	panic("宕机")
	
}
