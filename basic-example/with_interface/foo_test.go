package with_interface_test

import (
	"errors"
	"testing"
	
	"github.com/bygui86/go-testing/basic-example/with_interface"
)

// --- Success ---

type OkClientMock struct {
	GetDataReturn string
}

func (mc *OkClientMock) GetData() (string, error) {
	return mc.GetDataReturn, nil
}

func TestController_Success(t *testing.T) {
	err := with_interface.Controller(&OkClientMock{"data"})
	if err != nil {
		t.FailNow()
	}
}

// --- Failure ---

type FailingClientMock struct{}

func (fc *FailingClientMock) GetData() (string, error) {
	return "", errors.New("oh no")
}

// test failure of GetData()
func TestController_MockFailure(t *testing.T) {
	err := with_interface.Controller(&FailingClientMock{})
	if err == nil {
		t.FailNow()
	}
}

// test unexpected data returned from GetData()
func TestController_DataFailure(t *testing.T) {
	err := with_interface.Controller(&OkClientMock{"other-data"})
	if err == nil {
		t.FailNow()
	}
}
