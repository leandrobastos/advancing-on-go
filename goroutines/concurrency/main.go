package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	go runProcess("P1", 50)
	go runProcess("P2", 30)
	go runProcess("P3", 40)
	go runProcess("P4", 20)

	var s string
	fmt.Scanln(&s)
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		fmt.Println(name, " ->", i)

		t := time.Duration(rand.Intn(255))

		time.Sleep(time.Millisecond * t)
	}
}
