package main

import "fmt"

func naruto() func() {
	var count int
	count = 1
	return func() {
		var strPad = ""
		for i := 0; i < count; i++ {
			strPad += "!"
		}
		if count%2 != 0 {
			fmt.Println("naruto " + strPad)
		} else {
			fmt.Println("sasuke " + strPad)
		}
		count++
	}
}

func main() {
	//var test = naruto()
}
