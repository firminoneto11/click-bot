package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

func main() {
	for {
		robotgo.Click("left", true)
		fmt.Println("Clicked!")
		time.Sleep(time.Second * 2)
	}
}
