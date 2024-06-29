package factory_method

import "fmt"

type ChickenBurger struct {
	Name string
}

func (b *ChickenBurger) Prepare() {
	fmt.Printf("Preparing %s\n", b.Name)
}

type ChickenBurgerFactory struct{}

func (cbf *ChickenBurgerFactory) CreateBurger(name string) Burger {
	return &ChickenBurger{
		Name: name,
	}
}
