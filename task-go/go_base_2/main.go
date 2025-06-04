package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
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

	// Goroutine作业2
	taskScheduler()

	// 面向对象作业1、2
	obj()

	// channel作业1
	channel1()

	// channel作业2
	channel2()

	// 锁机制作业1
	lock1()

	// 锁机制作业2
	lock2()
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

	// Channel

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

// 任务调度器
func taskScheduler() {
	type Task struct {
		Name string
		Func func()
	}

	var wg sync.WaitGroup

	method := func(tasks []Task) {
		wg.Add(len(tasks))

		for _, task := range tasks {
			go func(t Task) {
				defer wg.Done() // 确保任务完成时调用 Done

				start := time.Now()
				fmt.Printf("开始执行任务: %s\n", t.Name)
				t.Func()
				duration := time.Since(start)
				fmt.Printf("任务 %s 执行时间: %v\n", t.Name, duration)
			}(task)
		}
	}
	tasks := []Task{
		{
			Name: "任务1",
			Func: func() {
				time.Sleep(time.Second * 2)
			},
		},
		{
			Name: "任务2",
			Func: func() {
				time.Sleep(time.Second * 3)
			},
		},
	}
	method(tasks) // 启动任务调度器
	wg.Wait()
}

func obj() {
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

func channel1() {
	method := func() {
		var ch = make(chan int)
		go func() {
			for i := 1; i <= 10; i++ {
				ch <- i // 将整数发送到通道
			}
			fmt.Println("发送完毕，关闭通道")
			close(ch) // 关闭通道
		}()

		go func() {
			for {
				select {
				case num, ok := <-ch:
					if ok {
						fmt.Println("接收到的整数:", num) // 从通道接收整数并打印
					} else {
						fmt.Println("通道已关闭")
						return
					}
				default:
					fmt.Println("等待接收数据...")
					time.Sleep(100 * time.Millisecond) // 避免忙等待
				}
			}
		}()
	}

	method()
	time.Sleep(2 * time.Second) // 确保协程有时间执行
}

func channel2() {
	method := func() {
		ch := make(chan int, 10) // 创建一个缓冲通道，容量为10

		go func() {
			for i := 1; i <= 100; i++ {
				ch <- i // 生产者将整数发送到通道
				fmt.Println("生产者发送:", i)
			}
			close(ch) // 关闭通道
		}()

		go func() {
			for num := range ch { // 从通道接收整数
				fmt.Println("消费者接收到:", num)
			}
			fmt.Println("消费者完成接收")
		}()
	}

	method()
	time.Sleep(5 * time.Second) // 确保协程有时间执行
}

type Counter struct {
	mu    sync.Mutex // 互斥锁
	count int
}

func (c *Counter) Increment() {
	c.mu.Lock()         // 上锁
	defer c.mu.Unlock() // 确保在函数结束时解锁
	c.count++           // 递增计数器
}

// 锁机制1
func lock1() {
	counter := &Counter{}

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				counter.Increment() // 每个协程递增计数器
			}
		}()
	}

	time.Sleep(2 * time.Second)            // 等待所有协程完成
	fmt.Println("计数器的最终值:", counter.count) // 输出计数器的值
}

// 锁机制2
func lock2() {
	var counter uint64

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				// 使用原子操作递增计数器
				atomic.AddUint64(&counter, 1)
			}
		}()
	}
	time.Sleep(2 * time.Second)        // 等待所有协程完成
	fmt.Println("无锁计数器的最终值:", counter) // 输出计数器的值
}
