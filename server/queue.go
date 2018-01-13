package server

import (
	"sync/atomic"
	"unsafe"
)

type node struct {
	T    interface{}
	Next unsafe.Pointer
}

type queue struct {
	Front, Rear unsafe.Pointer
}

func newQueue() *queue {
	node := &node{}
	return &queue{unsafe.Pointer(node), unsafe.Pointer(node)}
}

func (q *queue) Enqueue(T interface{}) {
	nElem := &node{
		T: T,
	}
	for {
		Rear := q.Rear
		Next := (*node)(Rear).Next

		if Rear != q.Rear {
			continue
		}

		if Next == nil {
			if atomic.CompareAndSwapPointer(&(*node)(q.Rear).Next, Next, unsafe.Pointer(nElem)) {
				atomic.CompareAndSwapPointer(&q.Rear, Rear, unsafe.Pointer(nElem))
				break
			}
		} else {
			atomic.CompareAndSwapPointer(&q.Rear, Rear, Next)
		}
	}
}

func (q *queue) Dequeue() (T interface{}) {
	for {
		Front, Rear := q.Front, q.Rear
		Next := (*node)(Front).Next

		if Front != q.Front {
			continue
		}

		if Front == Rear {
			if Next == nil {
				return nil
			}
			atomic.CompareAndSwapPointer(&q.Rear, Rear, Next)
		} else {
			if atomic.CompareAndSwapPointer(&q.Front, Front, Next) {
				return (*node)(Next).T
			}
		}

	}
}
