package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

// força o Go a utilizar o número máximo de CPUs
// Obs.: a partir da versão 1.5 do Go isso não é necessário
// func init() {
// 	runtime.GOMAXPROCS(runtime.NumCPU())
// }

func main() {
	// fmt.Println(runtime.NumCPU()) // Mostra a quantidade de CPUs
	waitGroup.Add(2)
	go runProcess("P1", 20)
	go runProcess("P2", 20)
	go runProcess("P3", 20)
	go runProcess("P4", 20)
	waitGroup.Wait()
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		fmt.Println(name, " ->", i)

		t := time.Duration(rand.Intn(255))

		time.Sleep(time.Millisecond * t)
	}
	waitGroup.Done()
}
