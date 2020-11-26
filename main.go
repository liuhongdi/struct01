package main

import "fmt"

//定义struct
type student struct {
	age int8
	name string
}
//指针接收器
func(s *student) modByPointer(){
	s.age =18
	s.name = "老刘他人不错";
}
//非指针接收器
func(s student) modNotByPointer(){
	s.age = 28
	s.name = "老刘他人实在";
}
//非指针接收器,返回副本
func(s student) modReturn() student{
	s.age = 38
	s.name = "老刘是个码农"
	return s
}

func main(){
	student:=new(student)
	fmt.Println(student)

	student.modByPointer()
	fmt.Println(student)

	student.modNotByPointer()
	fmt.Println(student)

	student2:=student.modReturn()
	fmt.Println(student2)
}