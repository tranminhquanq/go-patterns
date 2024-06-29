package factory_method

import "fmt"

type VeggieBurger struct {
	Name    string
	IsVegan bool
}

func (b *VeggieBurger) GetName() string {
	return b.Name
}
func (b *VeggieBurger) SetName(name string) {
	b.Name = name
}
func (b *VeggieBurger) Prepare() {
	fmt.Printf("Preparing %s\n", b.Name)
}

type VeggieBurgerFactory struct{}

func (vbf *VeggieBurgerFactory) CreateBurger(
	name string,
) Burger {
	return &VeggieBurger{
		IsVegan: true,
		Name:    name,
	}
}
