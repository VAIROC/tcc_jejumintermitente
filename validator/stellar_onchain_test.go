// +build !local

package validator

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStellarValidateAddress(t *testing.T) {
	client := StellarClient{Endpoint: os.Getenv("AVA_STELLAR_ENDPOINT")}
	validator :