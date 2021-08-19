package main

import "fmt"

func main() {
	webapp, err := BuildApp()
	if err != nil {
		fmt.Println("Failed to start app: ", err)
	}
	webapp.Run()
}
