package main

import (
	"fmt"
	"strings"
)

func main()  {
	names := []string{
		"DoG",
		"cat",
		"she3P",
		"SNAKE",
	}
	new_animals := applyTransform(names, strings.ToUpper)
	fmt.Println(new_animals)
	
}


func  applyTransform(vals []string, action converter ) [] string {
    for index, val := range vals {
		vals[index] = action(val)
	}
	return vals
} 

type converter func(string) string
