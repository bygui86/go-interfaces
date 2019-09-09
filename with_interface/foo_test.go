package with_interface_test

import (
	"errors"
	"testing"

	"github.com/bygui86/go-interfaces/with_interface"
)

// --- Success ---

type MockClient struct {
	GetDataReturn string
}

func (mc MockClient) GetData() (string, error) {
	return mc.GetDataReturn, nil
}

func TestController_Success(t *testing.T) {
	err := with_interface.Controller(MockClient{"data"})
	if err != nil {
		t.FailNow()
	}
}

// --- Error ---

type FailingClient struct{}

func (fc FailingClient) GetData() (string, error) {
	return "", errors.New("oh no")
}

func TestController_Failure(t *testing.T) {
	// test failure of GetData()
	err := with_interface.Controller(FailingClient{})
	if err == nil {
		t.FailNow()
	}
	// test unexpected data returned from GetData()
	err = with_interface.Controller(MockClient{"not data"})
	if err == nil {
		t.FailNow()
	}
}
