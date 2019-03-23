package main

import (
	"fmt"
)

func getPrefix(name string) (prefix string){
	var prefixMap map[string]string
	prefixMap = make(map[string]string)

	prefixMap["Bob"] = "Mr "
	prefixMap["Joe"] = "Dr "
	prefixMap["Mary"] = "Mrs "

	return prefixMap[name]
}

func main () {
	fmt.Println(getPrefix("Mary"));
}