package main

import(
	"strings"
	"fmt"
)

func greet(name string) string {
	clear_name := strings.TrimSpace(name)
	if clear_name == "" {
		return "Hello, World!"
	}
	return fmt.Sprintf("Hello, %s!", clear_name)
}
