package main

import (
	"github.com/bygui86/go-testing/basic-example/external"
	"github.com/bygui86/go-testing/basic-example/no_interface"
	"github.com/bygui86/go-testing/basic-example/with_interface"
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
