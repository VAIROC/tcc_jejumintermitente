// +build !local

package validator

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIOSTValidateAddress(t *testing.T) {
	client := IOSTClient{Endpoint