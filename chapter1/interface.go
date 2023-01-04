package ods

type Queue interface {
	Add(interface{}) interface{}
	Remove() interface{}
}

type Stack interface {
	Push(interface{}) interface{}
	Pop() interface{}
}

type List interface {
	Size() int
	Get(int) interface{}
	Set(int, interface{}) interface{}
	Add(int, interface{})
	Remove(int) interface{}
}

type USet interface {
	Size() int
	Add(interface{}) bool
	Remove(interface{}) interface{}
	Find(interface{}) interface{}
}

type SSet interface {
	Size() int
	Add(interface{}) bool
	Remove(interface{}) interface{}
	Find(interface{}) interface{}
}
