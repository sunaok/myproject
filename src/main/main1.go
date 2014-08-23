package main

import (
	"errors"
	f "fmt"
	"log"
	"os"
	//"github.com/sunaok/gosample"
)

func sum(i, j int) int {
	return i + j
}
func swap(i, j int) (int, int) {
	return j, i
}

func main() {
	file, err := os.Open("hello.go")
	if err != nil {
		f.Println(errors.New("Error"))
		log.Println("Error log")
	} else {
		file.Close()
	}
	f.Println("Error Sample")
	//	f.Println(sum(1,2))
	//	x, y := 1, 2
	//	x, y = swap(x,y)
	//	f.Println(x,y)
	//f.Print(gosample.Message)
	//a,b := 110,100
	//	if a > b {
	//		f.Print("a is larger than b")
	//	}else if a < b {
	//		f.Print("a is smaller than b")
	//	}else{
	//		f.Print("a equals b")
	////	}
	//	for i := 0;i < 10 ; i++{
	//		f.Println(i)
	//	}
	//	n := 0
	//	for n < 100{
	//		f.Println(n)
	//		n++
	//	}
	//
	//	n:=0
	//	for n<100{
	//	switch n{
	//	case 15:
	//		f.Println("Fizzbuzz")
	//	case 5,10:
	//		f.Println("Buzz")
	//	case 3,6,9:
	//		f.Println("Fizz")
	//	default:
	//		f.Println(n)
	//	}
	//		n++
	//	}

}
