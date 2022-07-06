package main

import(
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello from producer")
		time.Sleep(time.Second * 1)
	}
}