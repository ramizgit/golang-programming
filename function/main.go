package main

import "fmt"

func main() {

	rectangle := Rectangle{4, 5}

	//via function
	fmt.Println("Rectangle area via function : ", getRectangleAreaFunc(rectangle))

	//via method
	fmt.Println("Rectangle area via method : ", rectangle.getArea())
}

type Rectangle struct {
	length int
	bread  int
}

//function - this exists independently
func getRectangleAreaFunc(rectangle Rectangle) int {
	return rectangle.length * rectangle.bread
}

//method
func (rectangle Rectangle) getArea() int {
	return rectangle.length * rectangle.bread
}
