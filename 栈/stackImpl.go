package stack

import (
	"fmt"
	array "myGolang/数据结构/数组"
)

type StackImpl struct {
	arr *array.ArrayImpl
}

// NewStack 构造函数
func NewStack(cap int) *StackImpl {
	return &StackImpl{
		arr: array.NewArray(cap),
	}
}

// Getsize 获取栈size
func (s *StackImpl) GetSize() int {
	return s.arr.GetSize()
}

// IsEmpty 栈是否空
func (s *StackImpl) IsEmpty() bool {
	return s.arr.IsEmpty()
}

// Push 压栈
func (s *StackImpl) Push(v interface{}) {
	s.arr.AddLast(v)
}

// Pop 出栈
func (s *StackImpl) Pop() interface{} {
	return s.arr.RemoveLast()
}

// Peek 查看最后一个
func (s *StackImpl) Peek() interface{} {
	return s.arr.GetLast()
}

func (s *StackImpl) ToString() {
	size := s.arr.GetSize()
	length := s.arr.GetLength()
	capacity := s.arr.GetCapacity()
	fmt.Printf("栈：size=%d , len=%d , cap=%d \n", size, length, capacity)
	fmt.Print("[")
	for i := 0; i < size; i++ {
		fmt.Print(s.arr.Get(i))
		fmt.Print(" , ")
	}
	fmt.Println("] top")
}
