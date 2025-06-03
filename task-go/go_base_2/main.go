package main

import (
	"fmt"
	"sync"
)

func main() {
	// 指针作业1
	a := 2
	pointerBase1(&a)
	fmt.Println(a)

	// 指针作业2
	b := []int{1, 2, 3}
	pointerBase2(b)
	fmt.Println(b)

	// Goroutine作业1
	goroutine1()
}

// 编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
func pointerBase1(poin *int) {
	*poin *= 10
}

// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
func pointerBase2(poin []int) {
	for i := 0; i < len(poin); i++ {
		if poin[i] > 0 {
			poin[i] *= 2
		}
	}
}

func goroutine1() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan int)
	// 奇数
	go func() {
		defer wg.Done()
		oddNum := []int{}
		for i := 1; i <= 10; i++ {
			val, _ := <-ch
			if val > 0 && val%2 == 0 {
				oddNum = append(oddNum, val)
			}
		}
		fmt.Println("奇数部份:", oddNum)
	}()
	// 偶数
	go func() {
		defer wg.Done()
		evnNum := []int{}
		for i := 1; i <= 10; i++ {
			if i%2 == 1 {
				evnNum = append(evnNum, i)
			}
			ch <- i
		}
		fmt.Println("偶数部份:", evnNum)
	}()
	wg.Wait()

	// 面向对象 1
	// 先打印结构体的值
	var rectangle Rectangle
	fmt.Println("Rectangle.a:", rectangle.a)
	fmt.Println("Rectangle.b:", rectangle.b)
	var circle Circle
	fmt.Println("Circle.c:", circle.c)
	fmt.Println("Circle.d:", circle.d)
	rectangle = Rectangle{a: 1, b: "2"}

	// 执行方法
	rectangle.Area()
	circle.Perimeter()
	// 验证接口是否执行，值修改了就是执行了
	fmt.Println("new Rectangle.a:", rectangle.a)
	fmt.Println("new Rectangle.b:", rectangle.b)
	fmt.Println("new Circle.c:", circle.c)
	fmt.Println("new Circle.d:", circle.d)

	// 面向对象2
	emplInfo := Employee{Person{"JXD", 18}, 2131231}
	emplInfo.PrintInfo()
}

type Shape interface {
	Area()
	Perimeter()
}

func (rectangle *Rectangle) Area() {
	rectangle.a = 223
	rectangle.b = "test,test"
}

func (circle *Circle) Perimeter() {
	circle.c = 123
	circle.d = []int{123, 234, 345}
}

type Rectangle struct {
	a int
	b string
}

func (rectangle Rectangle) getRectangleA() int {
	return rectangle.a
}

func (rectangle Rectangle) getRectangleB() string {
	return rectangle.b
}

type Circle struct {
	c byte
	d []int
}

func (circle Circle) getCircleC() byte {
	return circle.c
}

func (circle Circle) getCircleD() []int {
	return circle.d
}

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	psn        Person
	employeeID int
}

func (e Employee) PrintInfo() {
	fmt.Println("员工姓名:", e.psn.Name)
	fmt.Println("员工年龄:", e.psn.Age)
	fmt.Println("员工ID:", e.employeeID)
}
