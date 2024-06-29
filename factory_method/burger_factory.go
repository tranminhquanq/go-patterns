package factory_method

type Burger interface {
	Prepare()
}
type BurgerFactory interface {
	CreateBurger(name string) Burger
}

type BurgerStore struct{}

func (bs *BurgerStore) BurgerFactory(burgerType string) BurgerFactory {
	switch burgerType {
	case "cheese":
		return &CheeseBurgerFactory{}
	case "veggie":
		return &VeggieBurgerFactory{}
	case "chicken":
		return &ChickenBurgerFactory{}
	default:
		return nil
	}
}
