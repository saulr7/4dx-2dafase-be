package main

import (
	"fmt"

	"./server"
)

func main() {
	fmt.Println("Corriendo")
	server.Serve()
}
