package main

func main() {

}

type rectangle struct {
	width  int
	height int
}

func (rec rectangle) calculateArea() int {
	return rec.width * rec.height
}

func (rec rectangle) calculatePerimeter() int {
	return rec.width * rec.height
}
