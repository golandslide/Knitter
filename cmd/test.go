package main

import (
	"Knitter/pkg/common"
	"fmt"
)

func main() {
	fmt.Println("Hi! Robot, I just wanna test Travis CI")
	if err := common.FooBarInvocation("true"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hello World")
	}
}
