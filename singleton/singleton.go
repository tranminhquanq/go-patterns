package singleton

import (
	"fmt"
	"sync"
)

type Single struct{}

var instance *Single
var lock = &sync.Mutex{}

func GetInstance() *Single {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		if instance == nil {
			instance = &Single{}
			fmt.Println("Returning new instance")
		} else {
			fmt.Println("Returning existing instance")
		}
	} else {
		fmt.Println("Returning existing instance")
	}

	return instance
}
