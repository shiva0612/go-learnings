package main

import (
	"fmt"
	"sync"
)

// define a struct for the object to be pooled
type PooledObject struct {
	// define any fields you need for your object
	name string
}

// define the pool struct
type ObjectPool struct {
	pool    chan *PooledObject
	maxSize int
	lock    sync.Mutex
}

// define the constructor function for creating new PooledObject instances
func NewPooledObject(name string) *PooledObject {
	return &PooledObject{name: name}
}

// define a function to create a new object pool with a given maximum size
func NewObjectPool(maxSize int) *ObjectPool {
	p := &ObjectPool{
		pool:    make(chan *PooledObject, maxSize),
		maxSize: maxSize,
	}
	for i := 0; i < cap(p.pool); i++ {
		p.pool <- NewPooledObject("default")
	}
	return p
}

// define a function to get an object from the pool
func (p *ObjectPool) Get() (*PooledObject, error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	select {
	case obj := <-p.pool:
		return obj, nil
	default:
		return nil, fmt.Errorf("ObjectPool: no more objects available")
	}
}

// define a function to return an object to the pool
func (p *ObjectPool) Release(obj *PooledObject) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	if len(p.pool) >= p.maxSize {
		return fmt.Errorf("ObjectPool: pool is already full")
	}

	p.pool <- obj
	return nil
}

func main() {
	// create a new object pool with a maximum size of 2
	pool := NewObjectPool(2)

	// get an object from the pool
	obj1, err := pool.Get()
	if err != nil {
		fmt.Println(err)
		return
	}

	// use the object
	fmt.Println("Object name:", obj1.name)

	// release the object back to the pool
	if err := pool.Release(obj1); err != nil {
		fmt.Println(err)
		return
	}

	// get another object from the pool
	obj2, err := pool.Get()
	if err != nil {
		fmt.Println(err)
		return
	}

	// use the second object
	fmt.Println("Object name:", obj2.name)

	// try to get a third object from the pool (should fail)
	obj3, err := pool.Get()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Object name:", obj3.name)
	}
}
