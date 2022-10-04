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
	validator := &Stellar{Client: &client}

	var cases = map[string]*Result{
		"GAI3GJ2Q3B35AOZJ36C4ANE3HSS4NK7WI6DNO4ZSHRAX6NG7BMX6VJER": {Success, true, Normal, "