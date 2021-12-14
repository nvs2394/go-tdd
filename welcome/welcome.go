package main

import "fmt"

func Welcome(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Print(Welcome("Son"))
}
