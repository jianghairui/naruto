package naruto

import "fmt"

func Naruto() func() {
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

func Hello() string {
	return "hello world!!!"
}
