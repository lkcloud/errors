package errors_test

/*
WARNING - changing the line numbers in this file will break the
examples.
*/

import (
	"fmt"

	errs "github.com/lkcloud/errors"
)

const (
	// Error codes below 1000 are reserved future use by the
	// "github.com/lkcloud/errors" package.
	ConfigurationNotValid errs.Code = iota + 1000
)

func init() {
	errs.Codes[ConfigurationNotValid] = errs.ErrCode{
		Ext:  "Configuration not valid",
		Int:  "the configuration is invalid",
		HTTP: 500,
	}
}

func loadConfig() error {
	err := decodeConfig()
	return errs.Wrap(err, ConfigurationNotValid, "service configuration could not be loaded")
}

func decodeConfig() error {
	err := readConfig()
	return errs.Wrap(err, errs.ErrInvalidJSON, "could not decode configuration data")
}

func readConfig() error {
	err := fmt.Errorf("read: end of input")
	return errs.Wrap(err, errs.ErrUnknown, "could not read configuration file")
}

func someWork() error {
	return fmt.Errorf("failed to do work")
}
