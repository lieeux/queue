package main

import (
	"fmt"
	"os"
	"queue/tools"
)

//type Queue struct { // 使用一个结构体管理队列
//	maxSize int
//	array   [5]int //模拟队列
//	head    int    //指向队首
//	tail    int    //指向队尾
//}
//
//func (q *Queue) IsFull() (b bool) { //判断队列是否为满
//	return (q.tail+1)%q.maxSize == q.head
//}
//
//func (q *Queue) IsEmpty() (b bool) { //判断队列是否为空
//	return q.head == q.tail
//}
//
//func (q *Queue) Push(val int) (err error) { // 添加数据
//	if q.IsFull() {
//		return errors.New("queue full")
//	}
//	q.array[q.tail] = val
//	q.tail = (q.tail + 1) % q.maxSize
//	return
//}
//
//func (q *Queue) Pop() (val int, err error) { // 取数据
//	if q.IsEmpty() {
//		return 404, errors.New("queue empty")
//	}
//	val = q.array[q.head]
//	q.head = (q.head + 1) % q.maxSize
//	return val, err
//}
//
//func (q *Queue) ShowQueue() { // 显示队列
//	for i := 0; i < (q.tail+q.maxSize-q.head)%q.maxSize; i++ {
//		fmt.Printf("array[%d]=%d\t", (q.head+i)%q.maxSize, q.array[(q.head+i)%q.maxSize])
//	}
//	fmt.Println()
//}

func main() {
	queue := &tools.Queue{
		MaxSize: 5,
		Head:    0,
		Tail:    0,
	}

	var key string
	var val int

	for {
		fmt.Println("命令：add\tget\tshow\texit")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入要添加的数据")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("add success")
			}
		case "get":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println(val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}

}
