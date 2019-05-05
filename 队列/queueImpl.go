package queue

import (
	"fmt"
	array "myGolang/数据结构/数组"
)

type ArrayQueue struct {
	arr *array.ArrayImpl
}

func NewQueue(cap int) *ArrayQueue {
	return &ArrayQueue{
		arr: array.NewArray(cap),
	}
}

// GetSize 获取队列size
func (a *ArrayQueue) GetSize() int {
	return a.arr.GetSize()
}

// IsEmpty 队列是否空
func (a *ArrayQueue) IsEmpty() bool {
	return a.arr.IsEmpty()
}

// Enqueue 入队
func (a *ArrayQueue) Enqueue(v interface{}) {
	a.arr.AddLast(v)
}

//Dequeue 出队
func (a *ArrayQueue) Dequeue() interface{} {
	return a.arr.RemoveFirst()
}

// GetFront 查看队首元素
func (a *ArrayQueue) GetFront() interface{} {
	return a.arr.GetFirst()
}

func (a *ArrayQueue) ToString() {
	size := a.arr.GetSize()
	length := a.arr.GetLength()
	capacity := a.arr.GetCapacity()
	fmt.Printf("queue ：size=%d , len=%d , cap=%d \n", size, length, capacity)
	fmt.Print(" front [")
	for i := 0; i < size; i++ {
		fmt.Print(a.arr.Get(i))
		fmt.Print(" , ")
	}
	fmt.Println("] tail")
}
