package array

import "fmt"

func GetFruits() {
	var fruits [10]string

	fruits[0] = "apple"
	fruits[1] = "orange"
	fruits[3] = "banana"

	fmt.Println("fuirts : ", fruits)
}
