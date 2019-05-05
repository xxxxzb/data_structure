package queue

type Queue interface {
	Enqueue(interface{})   //入队
	Dequeue() interface{}  //出队
	GetFront() interface{} //查看队首元素
	GetSize() int
	IsEmpty() bool
}
