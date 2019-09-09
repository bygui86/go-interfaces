package with_interface

import (
	"errors"
)

type IExternalClient interface {
	GetData() (string, error)
}

func Controller(externalClient IExternalClient) error {
	fromExternalAPI, err := externalClient.GetData()
	if err != nil {
		return err
	}
	// do some things based on data from external API
	if fromExternalAPI != "data" {
		return errors.New("unexpected data")
	}
	return nil
}
