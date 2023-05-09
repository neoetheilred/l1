package main

import "fmt"

func main() {
	action := Action{}
	action.Annoy()
	action.p = 10
	action.ID = 200
	action.Age = 31
	action.Name = "Tom"
	action.Troop = "uu"
	fmt.Println(action)       // {{200 31 Tom 10} uu}
	fmt.Println(action.Human) // {200 31 Tom 10}
}

type Human struct {
	ID   int
	Age  int
	Name string
	p    int
}

type Action struct {
	Human // This allows us to access Human
	Troop string
}

// Method of Human
func (h *Human) Annoy() {
	fmt.Println("AnnoyingAnnoyingAnnoyingAnnoyingAnnoyingAnnoyingAnnoyingAnnoyingAnnoyingAnnoyingAnnoying")
}
