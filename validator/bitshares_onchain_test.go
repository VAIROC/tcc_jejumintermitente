// +build !local

package validator

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitsharesValidateAddress(t *testing.T) {
	client := BitsharesClient{Endpoint: os.Getenv("AVA_BITSHARES_ENDPOINT")}
	validator := &Bitshares{Client: &client}

	var cases =