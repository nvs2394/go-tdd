package main

import "fmt"

const englishWelcomPrefix = "Hello, "
const germanWelcomPrefix = "Hallo, "

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case "German":
		prefix = germanWelcomPrefix
	default:
		prefix = englishWelcomPrefix
	}
	return prefix
}

func Welcome(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(lang) + name
}

func main() {
	fmt.Print(Welcome("Son", ""))
}
