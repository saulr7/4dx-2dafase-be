package main

import (
	"fmt"

	"./server"
)

func main() {
	server.Serve()
	fmt.Println("Corriendo...")
}
