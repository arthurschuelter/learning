package main

import (
	"fmt"
)

type DeleteFunction interface {
	delete(c *Cache)
}

type Cache struct {
	storage        map[string]string
	deleteFunction DeleteFunction
	capacity       int
	maxCapacity    int
}

func initCache(del DeleteFunction) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:        storage,
		deleteFunction: del,
		capacity:       0,
		maxCapacity:    2,
	}
}

func (c *Cache) addValue(key, value string) {
	if c.capacity >= c.maxCapacity {
		c.deleteFunction.delete(c)
		return
	}

	c.capacity++
	c.storage[key] = value
}

func (c *Cache) setDeleteFunction(df DeleteFunction) {
	c.deleteFunction = df
}

func (c *Cache) print() {
	fmt.Println(c.storage)
	fmt.Printf("Capacity: %d/%d\n", c.capacity, c.maxCapacity)
}

func main() {
	fmt.Println("Strategy Pattern")

	fifo := &Fifo{}

	cache := initCache(fifo)
	cache.addValue("a", "b")
	cache.addValue("b", "c")
	cache.print()

	cache.addValue("c", "d")

	lifo := &Lifo{}
	cache.setDeleteFunction(lifo)
	cache.addValue("c", "d")

}
