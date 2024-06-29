package main

import (
	"github.com/tranminhquanq/go-patterns/factory_method"
	"github.com/tranminhquanq/go-patterns/singleton"
)

func main() {
	// // Singleton
	for i := 0; i < 10; i++ {
		go singleton.GetInstance()
	}

	// Factory Method
	burgerStore := factory_method.BurgerStore{}

	cheeseBurgerFactory := burgerStore.BurgerFactory("cheese")
	cheddarBurger := cheeseBurgerFactory.CreateBurger("Cheddar Burger")
	provoloneBurger := cheeseBurgerFactory.CreateBurger("Burger Burger")
	cheddarBurger.Prepare()
	provoloneBurger.Prepare()

	veggieBurgerFactory := burgerStore.BurgerFactory("veggie")
	quinoaBurger := veggieBurgerFactory.CreateBurger("Quinoa Burger")
	falafelBurger := veggieBurgerFactory.CreateBurger("Falafel Burger")
	quinoaBurger.Prepare()
	falafelBurger.Prepare()
}
