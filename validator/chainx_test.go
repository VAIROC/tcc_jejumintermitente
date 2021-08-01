package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChainXValidateAddress(t *testing.T) {
	validator := &Cha