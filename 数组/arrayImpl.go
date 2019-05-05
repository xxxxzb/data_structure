package array

import "fmt"

// ArrayImpl 自定义的数组结构
type ArrayImpl struct {
	data []interface{} // 泛型数组
	size int           // 元素数量
}

// GetArray 带参构造函数
func NewArray(capacity int) *ArrayImpl {
	return &ArrayImpl{
		data: make([]interface{}, capacity),
		size: 0,
	}
}

// GetCapacity 获得数组容量
func (a *ArrayImpl) GetCapacity() int {
	return cap(a.data)
}

// GetLength 获得数组容量
func (a *ArrayImpl) GetLength() int {
	return len(a.data)
}

// GetSize 获得元素个数
func (a *ArrayImpl) GetSize() int {
	return a.size
}

// IsEmpty 查看数组是否为空
func (a *ArrayImpl) IsEmpty() bool {
	return a.size == 0
}

// 容量调整
// newCap 新数组容量
// 逻辑：声明新的数组，将原数组的值 copy 到新数组中
func (a *ArrayImpl) resize(newCap int) {
	newArray := make([]interface{}, newCap)
	for i := 0; i < a.size; i++ {
		newArray[i] = a.data[i]
	}
	a.data = newArray
}

// Find 查找元素返回第一个索引
func (a *ArrayImpl) Find(v interface{}) int {
	for i := 0; i < a.size; i++ {
		if a.data[i] == v {
			return i
		}
	}
	return -1
}

// FindAll 查找元素返回所有索引
func (a *ArrayImpl) FindAll(v interface{}) (res []int) {
	for i := 0; i < a.size; i++ {
		if a.data[i] == v {
			res = append(res, i)
		}
	}
	return
}

// Contains 查找是否存在元素
func (a *ArrayImpl) Contains(v interface{}) bool {
	for i := 0; i < a.size; i++ {
		if a.data[i] == v {
			return true
		}
	}
	return false
}

// Get by index
func (a *ArrayImpl) Get(index int) interface{} {
	if index < 0 || index > a.size-1 {
		panic("idnex 错误")
	}
	return a.data[index]
}
func (a *ArrayImpl) GetFirst() interface{} {
	return a.Get(0)
}
func (a *ArrayImpl) GetLast() interface{} {
	return a.Get(a.size - 1)
}

// Set by index
func (a *ArrayImpl) Set(index int, v interface{}) {
	if index < 0 || index > a.size-1 {
		panic("idnex 错误")
	}
	a.data[index] = v
}

// Add 添加元素
// 添加元素需要考虑扩容问题
func (a *ArrayImpl) Add(index int, v interface{}) {
	if index < 0 || index > a.GetCapacity() {
		panic("idnex 错误")
	}

	// 数组已满则扩容
	if a.size == a.GetCapacity() {
		a.resize(2 * a.size)
	}

	// 将插入的索引位置之后的元素后移，腾出插入位置
	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}
	a.data[index] = v
	a.size++
}
func (a *ArrayImpl) AddLast(v interface{}) {
	a.Add(a.size, v)
}

func (a *ArrayImpl) AddFirst(v interface{}) {
	a.Add(0, v)
}

// Remove 删除 + 缩容
func (a *ArrayImpl) Remove(index int) (removeElement interface{}) {
	if index < 0 || index > a.size {
		panic("idnex 错误")
	}
	removeElement = a.data[index]

	// 从 index 之后的元素，都向前移动一个位置
	for i := index + 1; i < a.size; i++ {
		a.data[i-1] = a.data[i]
	}
	a.size--
	// a.size == len(a.data)/4 防止扩缩容震荡
	// 考虑边界情况，不能 resize 为0
	if a.size == len(a.data)/4 && len(a.data)/2 != 0 {
		a.resize(len(a.data) / 2)
	}
	return
}
func (a *ArrayImpl) RemoveFirst() interface{} {
	return a.Remove(0)
}

func (a *ArrayImpl) RemoveLast() interface{} {
	return a.Remove(a.size - 1)
}

// ToString 打印
func (a *ArrayImpl) ToString() {
	fmt.Printf("数组：size=%d , len=%d , cap=%d \n", a.size, len(a.data), cap(a.data))
	fmt.Print("[")
	for i := 0; i < a.size; i++ {
		fmt.Print(a.data[i])
		fmt.Print(" , ")
	}
	fmt.Println("]")
}
