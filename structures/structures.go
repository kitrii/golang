package main

import "fmt"

func main() {
	/*
				type Circle struct {
			    x float64
			    y float64
			    r float64
			}
		type Circle struct {
			x, y, r float64
		}
	*/
	//structure initialization
	//create a new object of type Circle
	// b := Circle{x: 4, y: 5, r: 6}
	// var c Circle = Circle{8, 8, 8}
	// d := Circle{1, 1, 3}
	// e := new(Circle) //таким образом создается только пустой экземпляр

	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)
	ob := Circle{10, 20, 30}
	fmt.Println(Circle2(ob))

}

type Circle struct {
	x, y, r float64
}

func Circle2(object Circle) float64 {
	return object.x * object.y * object.r
}

// we can yse this func
func Circle1(x, y, r float64) float64 {
	return x * y * r

}
