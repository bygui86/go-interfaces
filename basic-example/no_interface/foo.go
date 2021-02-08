package no_interface

import (
	"errors"

	"github.com/bygui86/go-testing/basic-example/external"
)

func Controller() error {
	externalClient := external.NewClient()
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
