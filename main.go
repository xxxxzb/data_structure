package main

import (
	queue "myGolang/数据结构/队列"
)

func main() {
	// //==================数组
	// arr := array.NewArray(10)
	// for i := 0; i < 10; i++ {
	// 	arr.AddLast(i)
	// }
	// arr.ToString()

	// arr.Add(1, 100)
	// arr.ToString()

	// arr.AddFirst(-1)
	// arr.ToString()
	// //==================数组

	// //==================栈
	// stack := stack.NewStack(10)
	// for i := 0; i < 10; i++ {
	// 	stack.Push(i)
	// 	stack.ToString()
	// }

	// stack.Pop()
	// stack.ToString()
	// //==================栈

	//==================队列
	queue := queue.NewQueue(10)
	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
		// queue.ToString()

		if i%3 == 2 {
			queue.Dequeue()
			queue.ToString()
		}
	}
	//==================队列

}
