// +build !local

package validator

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIOSTValidateAddress(t *testing.T) {
	client := IOSTClient{Endpoint: os.Getenv("AVA_IOST_ENDPOINT")}
	validator := &IOST{Client: &client}

	var mainnetCases = map[string]*Result{
		"k