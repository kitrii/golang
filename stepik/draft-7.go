package main

import (
	"fmt"
	"strings"
)

var t interface {
	talk() string
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

func main() {

	m := t

	m = martian{}
	fmt.Println(m.talk())

	m = laser(9)
	fmt.Println(m.talk())

}
