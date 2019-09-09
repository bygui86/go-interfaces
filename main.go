package main

import (
	"github.com/bygui86/go-interfaces/external"
	"github.com/bygui86/go-interfaces/no_interface"
	"github.com/bygui86/go-interfaces/with_interface"
)

func main() {

	noErr := no_interface.Controller()
	if noErr != nil {
		panic(noErr)
	}

	withErr := with_interface.Controller(external.NewClient())
	if withErr != nil {
		panic(withErr)
	}
}
