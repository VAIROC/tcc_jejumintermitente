// +build !local

package validator

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOasisValidateAddress(t *testing.T) {
	client := RosettaClient{Endpoint: os.Getenv("AVA_OASIS