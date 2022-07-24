// +build !local

package validator

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOasisValidateAddress(t *testing.T) {
	client := RosettaClient{Endpoint: os.Getenv("AVA_OASIS_ENDPOINT")}
	validator := &Oasis{Client: &client}

	var cases = map[string]*Result{
		"oasis1qqs6ylpfurhf6gc9mw232fkmrt3d0673lyzc5xf2": {Success, true, Norm