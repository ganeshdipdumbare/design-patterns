package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type single struct {
}

var singleInstance *single

func getInstance() *single {

	once.Do(
		func() {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		})

	fmt.Println("Single instance already created.")

	return singleInstance
}
