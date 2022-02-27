package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var x int
	y := 7
	x = 8
	fmt.Println(x, y)
	myValue, err := strconv.ParseInt("7", 0, 64)
	if err != nil {
		fmt.Printf("%v/n", err)
	} else {
		fmt.Println(myValue)
	}

	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)

	s := []int{1, 2, 3, 4, 5}
	for index, value := range s {
		fmt.Println(index, value)
	}

	s = append(s, 6)

	for index, value := range s {
		fmt.Println(index, value)
	}
	// c := make(chan int)
	// go doSomething(c)
	// <-c

	g := 25
	fmt.Println(g)

	h := &g
	fmt.Println(h)
	fmt.Println(*h)
}

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Doing something")
	c <- 1
}
