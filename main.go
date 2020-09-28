package main

import "fmt"

func main()  {
	var lado uint64

	fmt.Scanln(&lado)

	area:= lado *lado

	fmt.Println(area)
}