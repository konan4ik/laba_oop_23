package utilities

import "fmt"

type Event[T any] struct {
	handlers []func(T)
}

func (e *Event[T]) Subscribe(h func(T)) {
	e.handlers = append(e.handlers, h)
}

func (e *Event[T]) Unsubscribe(h func(T)) {
	for i, v := range e.handlers {
		if fmt.Sprint("%p", v) == fmt.Sprint("%p", h) {
			e.handlers = append(e.handlers[:i], e.handlers[i+1:]...)
			return
		}
	}
}

func (e *Event[T]) Emit(data T) {
	for _, h := range e.handlers {
		h(data)
	}
}

func NewEvent[T any]() *Event[T] {
	return &Event[T]{}
}
