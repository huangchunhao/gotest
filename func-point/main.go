package main

import (
	"fmt"
	"github.com/huangchunhao/goutil"
)

type testint int
type testsint struct {
  sint testint
}

func (s *testsint) testdoubles() int{
	sv :=s.sint
	s.sint=sv*2
	ssss:="5656566"
	goutil.Pfv(&ssss)
	goutil.PfT(&ssss)
	fmt.Printf("1-的方法内sv的值%v\n",sv)
	fmt.Printf("1-的方法内s的值%v\n",*s)
	return 0
}

func (sw testsint) testdoublesw() int{
	sv :=sw.sint
	sw.sint=sv*2
	fmt.Printf("2-的方法内sv的值%v\n",sv)
	fmt.Printf("2-的方法内sw的值%v\n",sw)
	return 0
}

func main(){
	testsint1 :=new(testsint)
	testsint1.sint=2
	testsint1.testdoubles()
	fmt.Printf("1-的main内s的值%v\n",*testsint1)


	testsint2 :=new(testsint)
	testsint2.sint=3
	testsint2.testdoublesw()
	fmt.Printf("2-的main内sw的值%v\n",*testsint2)
}