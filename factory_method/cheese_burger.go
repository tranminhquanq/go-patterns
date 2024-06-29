package factory_method

import "fmt"

type CheeseBurger struct {
	Name             string
	AdditionalCheese bool
}

func (b *CheeseBurger) Prepare() {
	fmt.Printf("Preparing %s\n", b.Name)
}

type CheeseBurgerFactory struct{}

func (cbf *CheeseBurgerFactory) CreateBurger(
	name string,
) Burger {
	return &CheeseBurger{
		AdditionalCheese: true,
		Name:             name,
	}
}

func (b *CheeseBurger) GetName() string {
	return b.Name
}
