package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

var sandwichNum = 0
var sandwichName = ""

func main() {
	runtime.GOMAXPROCS(1)
	stringChan := make(chan string)

	for i := 0; i < 3; i++ {
		go getBread(stringChan)
		go addFilling(stringChan)
		go addSauce(stringChan)
		time.Sleep(time.Millisecond * 5000)
	}
}

func getBread(stringChan chan string) {
	sandwichNum++
	sandwichName = "Sandwich #" + strconv.Itoa(sandwichNum)
	fmt.Println("Getting bread and send for filling")
	stringChan <- sandwichName
	time.Sleep(time.Millisecond * 10)
}

func addFilling(stringChan chan string) {
	sandwich := <-stringChan
	fmt.Println("Add filling and send", sandwich, "for sauce")
	stringChan <- sandwichName
	time.Sleep(time.Millisecond * 10)
}

func addSauce(stringChan chan string) {
	sandwich := <-stringChan
	fmt.Println("Add sauce and ship", sandwich)
	stringChan <- sandwichName
	time.Sleep(time.Millisecond * 10)
}
