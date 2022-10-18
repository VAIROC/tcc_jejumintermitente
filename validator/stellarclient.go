package validator

import (
	"fmt"

	"github.com/LanfordCai/ava/httpclient"
)

// StellarClient ...
type StellarClient struct {
	Endpoint string
}

// GetAccount ...
func (c *StellarClient) GetAccount(addr string) (AddressType, error) {
	url := fmt.Sprintf("%s/accounts/%s", c.