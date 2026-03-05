package main

var c = make(chan int)
var a string

func f() {
	a = "hello, world" // 4
	<-c // 5
}

func main() {
	go f() // 1
	c <- 0 // 2
	print(a) // 3
}