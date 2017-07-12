package mypackage

import "fmt"

var i int = 0;

func main(){
	fmt.Println("worked")
}

func MyFunc(){
	i++
	fmt.Println("worked", i)
}