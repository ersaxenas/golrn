package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	fmt.Println("hello \n this works")
	fmt.Println('a')
	fmt.Println(math.Floor(10.3))
	fmt.Println(strings.Title("head first go"))
	// find out the type of any
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(42.9989))
	// variable declaration
	var quantity int
	var newname string
	quantity = 2
	fmt.Println(quantity)
	// assign multiple values
	quantity, newname = 5, "headfirst"
	fmt.Println(quantity, newname)
	// declare and assign value
	var bookname string = "head first go"
	fmt.Println(bookname)
	// short variable declaration with :=
	squantity := 100
	sbookname := "short head first go"
	fmt.Println(squantity)
	fmt.Println(sbookname)
	fmt.Println("format")
}
