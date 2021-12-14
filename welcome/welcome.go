package main

import "fmt"

const englishWelcomPrefix = "Hello, "

func Welcome(name string) string {
	if name == "" {
		name = "World"
	}
	return englishWelcomPrefix + name
}

func main() {
	fmt.Print(Welcome("Son"))
}
