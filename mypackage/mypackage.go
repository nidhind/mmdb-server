package mypackage

import "fmt"

var i int = 0;

func myMain(){
	fmt.Println("worked")
}

func MyFunc(){
	i++
	fmt.Println("worked", i)
	myMain()
}