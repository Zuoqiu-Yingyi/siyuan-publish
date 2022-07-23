package models

type model interface {
	Empty() bool
	One(id string) interface{}
}
