package array

// Array 数组的接口
type Array interface {
	// 自调用的方法
	resize() // 用于数组的扩、缩容

	// 基本方法
	GetCapacity() int // 获得数组容量
	GetSize() int     // 获得元素个数
	IsEmpty() bool    // 查看数组是否为空
	// 查找
	Find(interface{}) int      // 查找元素返回第一个索引
	FindAll(interface{}) []int // 查找元素返回所有索引
	Contains(interface{}) bool // 查找是否存在元素
	Get(int) interface{}
	GetFirst(int) interface{}
	GetLast(int) interface{}
	// 修改
	Set(int, interface{})
	// 添加
	Add(int, interface{}) // 插入元素
	AddLast(interface{})
	AddFirst(interface{})
	// 删除
	Remove(int) interface{}
	RemoveFirst() interface{}
	RemoveLast() interface{}
}
