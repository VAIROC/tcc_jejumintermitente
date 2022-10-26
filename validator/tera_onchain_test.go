// +build !local

package validator

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeraValidateAddress(t *testing.T) {
	client := TeraClient{Endpoint: os.Getenv("AVA_TERA_ENDPOINT")}
	validator := &Tera{Client: 