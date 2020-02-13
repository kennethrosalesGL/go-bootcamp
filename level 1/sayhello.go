package main

import (
	"fmt"
	"os"
)

func main()  {
	os_name :=  os.Args[1:][len(os.Args[1:])-1] 
	if len(os_name) > 0 { 
		fmt.Printf("Hello, %s", os_name)
	 } else {
		fmt.Println("What is your name?.")
		var name string
		fmt.Scanln(&name)
		fmt.Printf("Hello, %s", name)
	 }
}