package main

import "fmt"

type Fifo struct {
}

func (f *Fifo) delete(c *Cache) {
	fmt.Println("Deleting with fifo")
}

type Lifo struct {
}

func (l *Lifo) delete(c *Cache) {
	fmt.Println("Deleting with lifo")
}
