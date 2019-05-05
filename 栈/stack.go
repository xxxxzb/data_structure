package stack

type Stack interface {
	Getsize() int
	IsEmpty() bool
	Push(interface{})
	Pop() interface{}
	Peek() interface{} //查看最后一个
}
