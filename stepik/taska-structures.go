package main

import "fmt"

func main() {
	object := Game{true, 100, 0}
	fmt.Println(object)
	a := object.Shoot())
	fmt.Println(object.RideBike())

	// testStruct := *object

}

type Game struct {
	On    bool
	Ammo  int
	Power int
}

// }
func (g Game) Shoot() bool {
	var flag bool
	if g.On == false {
		flag = false
	} else if g.On == true && g.Ammo > 0 {
		g.Ammo -= 1
		flag = true
	}
	return flag
}

func (g Game) RideBike() bool {
	var flag bool
	if g.On == false {
		flag = false
	} else if g.On == true && g.Power > 0 {
		flag = true
	}
	return flag
}
