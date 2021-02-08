package no_interface_test

import (
	"testing"

	"github.com/bygui86/go-testing/basic-example/no_interface"
)

func TestController_Success(t *testing.T) {
	err := no_interface.Controller()
	if err != nil {
		t.FailNow()
	}
}

// WARN: this won't pass :(
// we want this to error but we can't get it to because we can't easily stub the external.Client struct
// func TestController_Error(t *testing.T) {
// 	err := no_interface.Controller()
// 	if err == nil {
// 		t.FailNow() // test will fail here
// 	}
// }
