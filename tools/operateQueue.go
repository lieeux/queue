package tools

import (
	"errors"
	"fmt"
)

type Queue struct { // 使用一个结构体管理队列
	MaxSize int
	array   [5]int //模拟队列
	Head    int    //指向队首
	Tail    int    //指向队尾
}

func (q *Queue) IsFull() (b bool) { //判断队列是否为满
	return (q.Tail+1)%q.MaxSize == q.Head
}

func (q *Queue) IsEmpty() (b bool) { //判断队列是否为空
	return q.Head == q.Tail
}

func (q *Queue) Push(val int) (err error) { // 添加数据
	if q.IsFull() {
		return errors.New("queue full")
	}
	q.array[q.Tail] = val
	q.Tail = (q.Tail + 1) % q.MaxSize
	return
}

func (q *Queue) Pop() (val int, err error) { // 取数据
	if q.IsEmpty() {
		return 404, errors.New("queue empty")
	}
	val = q.array[q.Head]
	q.Head = (q.Head + 1) % q.MaxSize
	return val, err
}

func (q *Queue) ShowQueue() { // 显示队列
	for i := 0; i < (q.Tail+q.MaxSize-q.Head)%q.MaxSize; i++ {
		fmt.Printf("array[%d]=%d\t", (q.Head+i)%q.MaxSize, q.array[(q.Head+i)%q.MaxSize])
	}
	fmt.Println()
}
