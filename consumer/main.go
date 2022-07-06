package main

import(
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello from consumer")
		time.Sleep(time.Second * 1)
	}
}