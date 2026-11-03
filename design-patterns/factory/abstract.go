package main

import "fmt"

type IComputer interface {
	setOS(os string)
	setPrice(price int)

	getOS() string
	getPrice() int

	print()
}

type Computer struct {
	os    string
	price int
}

func (c *Computer) setOS(os string) {
	c.os = os
}

func (c *Computer) setPrice(price int) {
	c.price = price
}

func (c *Computer) getOS() string {
	return c.os
}

func (c *Computer) getPrice() int {
	return c.price
}

func (c *Computer) print() {
	fmt.Println("Computer:", c.os, "Price: ", c.price)
}
