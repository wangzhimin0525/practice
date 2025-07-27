package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

func addTen(num *int) {
	*num += 10 // 解引用指针并修改值
}

func doubleSlice(slice []int) {
	for i := range slice {
		slice[i] *= 2
	}
}

func printOdds(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		fmt.Println("奇数", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printEvens(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		fmt.Println("偶数", i)
		time.Sleep(100 * time.Millisecond)
	}
}

// 定义任务类型 - 一个没有参数也没有返回值的函数
type Task func()

// 任务调度器
func runTasks(tasks []Task) {
	var wg sync.WaitGroup
	for i, task := range tasks {
		wg.Add(1)
		go func(taskIndex int, task Task) {
			defer wg.Done()
			start := time.Now()
			task()
			duration := time.Since(start)
			fmt.Printf("任务 %d 完成，耗时: %v\n", taskIndex+1, duration)
		}(i, task)
	}
	wg.Wait()
}

// Shape 接口定义
type Shape interface {
	Area() float64      // 计算面积
	Perimeter() float64 // 计算周长
}

// Rectangle 矩形结构体
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle 圆形结构体
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Person 结构体定义
type Person struct {
	Name string
	Age  int
}

// Employee 结构体通过组合 Person 来扩展
type Employee struct {
	Person
	EmployeeID string
}

// PrintInfo 方法属于 Employee 类型
func (e Employee) PrintInfo() {
	fmt.Printf("  姓名: %s\n", e.Name)
	fmt.Printf("  年龄: %d\n", e.Age)
	fmt.Printf("  员工ID: %s\n", e.EmployeeID)
}

// 生产者协程：生成1-10的数字并发送到通道
func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		fmt.Printf("生产者发送: %d\n", i)
		ch <- i // 发送数据到通道
	}
	close(ch)
}

// 消费者协程：从通道接收并打印数字
func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch { // 使用range持续接收直到通道关闭
		fmt.Printf("消费者接收: %d\n", num)
		time.Sleep(100 * time.Millisecond) // 模拟处理延迟
	}
	fmt.Println("通道已关闭，消费者结束")
}

// 生产者：生成1-100的数字并发送到缓冲通道
func producerWithBuffer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		ch <- i // 发送数据到缓冲通道
		fmt.Printf("生产: %d (通道长度: %d)\n", i, len(ch))
		time.Sleep(50 * time.Millisecond) // 模拟生产耗时
	}
	close(ch) // 发送完成后关闭通道
	fmt.Println("生产者完成生产")
}

// 消费者：从缓冲通道接收并打印数字
func consumerWithBuffer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch { // 使用range持续接收直到通道关闭
		fmt.Printf("消费: %d (通道长度: %d)\n", num, len(ch))
		time.Sleep(100 * time.Millisecond) // 模拟消费耗时
	}
	fmt.Println("消费者完成消费")
}

// 同步锁
type syncCounter struct {
	mu    sync.Mutex
	count int
}

func (c *syncCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *syncCounter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	//x := 5
	//fmt.Println("修改前的值:", x)
	//addTen(&x)
	//fmt.Println("修改后的值:", x)

	//nums := []int{1, 2, 3, 4, 5}
	//fmt.Println("修改前的切片:", nums)
	//doubleSlice(nums)
	//fmt.Println("修改后的切片:", nums)

	//var wg sync.WaitGroup
	//wg.Add(2)
	//go printOdds(&wg)
	//go printEvens(&wg)
	//wg.Wait()

	//tasks := []Task{
	//	func() {
	//		time.Sleep(1 * time.Second)
	//		fmt.Println("任务1执行中...")
	//	},
	//	func() {
	//		time.Sleep(2 * time.Second)
	//		fmt.Println("任务2执行中...")
	//	},
	//	func() {
	//		time.Sleep(1 * time.Second)
	//		fmt.Println("任务3执行中...")
	//	},
	//}
	//fmt.Println("开始执行任务...")
	//startTime := time.Now()
	//runTasks(tasks)
	//fmt.Printf("所有任务完成，总耗时: %v\n", time.Since(startTime))

	//// 创建矩形实例
	//rect := Rectangle{
	//	Width:  5.0,
	//	Height: 3.0,
	//}
	//// 创建圆形实例
	//circle := Circle{
	//	Radius: 4.0,
	//}
	//var s1 Shape = rect
	//fmt.Println("Rectangle Area:", s1.Area())
	//fmt.Println("Rectangle Perimeter:", s1.Perimeter())
	//var s2 Shape = circle
	//fmt.Println("Circle Area:", s2.Area())
	//fmt.Println("Circle Perimeter:", s2.Perimeter())

	//// 创建 Employee 实例
	//emp := Employee{
	//	Person: Person{
	//		Name: "zhangsan",
	//		Age:  3,
	//	},
	//	EmployeeID: "emp-1",
	//}
	//emp.PrintInfo()

	//var wg sync.WaitGroup
	//wg.Add(2)
	//ch := make(chan int)
	//go producer(ch, &wg)
	//go consumer(ch, &wg)
	//wg.Wait()

	//ch := make(chan int, 10)
	//var wg sync.WaitGroup
	//wg.Add(2)
	//go producerWithBuffer(ch, &wg)
	//go consumerWithBuffer(ch, &wg)
	//wg.Wait()
	//fmt.Println("所有任务完成")

	//counter := syncCounter{}
	//// 启动100个goroutine同时增加计数
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		for j := 0; j < 1000; j++ {
	//			counter.Increment()
	//		}
	//	}()
	//}
	//// 等待一段时间确保所有goroutine完成
	//time.Sleep(time.Second)
	//// 输出最终计数
	//fmt.Printf("Final count: %d\n", counter.GetCount())

	var counter int64 // 使用int64类型，因为atomic包支持的操作需要特定类型
	var wg sync.WaitGroup
	const numGoroutines = 10
	const incrementsPerGoroutine = 1000
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				atomic.AddInt64(&counter, 1) // 原子递增操作
			}
		}()
	}
	wg.Wait()
	fmt.Printf("最终计数器值: %d\n", counter)
}
