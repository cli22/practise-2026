package main

import "time"
var a string

func main() {
	hello2()
}

func f() {
	print(a) // 3
}

func hello() {
	a = "hello, world" // 1
	go f() // 2
	time.Sleep(time.Second)
}
func hello2() {
	go func() { a = "hello" }() // 1
	print(a) // 2
}