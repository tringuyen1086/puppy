package puppy

import (
	"fmt"

	"github.com/tringuyen1086/dog"
)

func Bark() string{
	return "Woof!"
}

func Barks() string{
	return "Woof! Woof! Woof!"
}

func BigBark() string{
	return dog.WhenGrowUp(Bark())
}

func BigBarks() string{
	return dog.WhenGrowUp(Barks())
}

func From11(){
	fmt.Println("version 1.1.0")
}

func From12(){
	fmt.Println("version 1.2.0")
}