package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age int
}

func (person Person) getAgeMinusTwo() int {
	return person.Age - 2
}


func main() {
	fmt.Println("Hello mwai")
	var name  = "Mwai"
	fmt.Printf("[ASSIGNED] %s\n", name)
	inferred := name
	fmt.Printf("[INFERRED] %s\n", inferred)

	numArr := []int { 1, 3, 4}
	numArr = append(numArr, 10)
	for i, v := range numArr {
		fmt.Printf("i=%d: %d\n", i, v)
	}

	/* PERSON */
	p := Person { "Mwai", 22 }
	fmt.Println("Age:", p.getAgeMinusTwo())

	fmt.Println("filtered:", filterIntArr(numArr, func(elem int) bool {
		return elem < 4
	}))

	fmt.Println("filtered:", filterIntArr(numArr, func(elem int) bool {
		return elem > 3
	}))

	fmt.Println(numArr)
	fmt.Println("map x + 1:")
	fmt.Println(mapIntArr(numArr, func(elem int) int {
		return elem + 1
	}))
}

func filterIntArr(arr []int, f func(int) bool ) []int {
	res := []int {}
	for i := range arr {
		if(f(arr[i])) {
			res = append(res, arr[i])
		}
	}
	return res
}

func mapIntArr(arr []int, f func(int) int ) []int {
	res := []int {}
	for i := range arr {
		res = append(res, f(arr[i]))
	}
	return res
}


