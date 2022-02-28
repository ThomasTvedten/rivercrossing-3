package main

import (
	"fmt"
)   "github.com/ThomasTvedten/rivercrossing-3/State"

func main() {
	fmt.Println(state.ViewState())

	state.Enterboat()
	fmt.Println(state.ViewState())

	state.Rivercross()
	fmt.Println(state.ViewState())
}
