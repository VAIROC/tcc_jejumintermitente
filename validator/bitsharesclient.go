package validator

import (
	"github.com/LanfordCai/ava/httpclient"
)

// BitsharesClient ...
type BitsharesClient struct {
	Endpoint string
}

// GetAccount ...
func (c *BitsharesClient) GetAccount(addr string) (AddressType, error) {
	reqBody := map[string]interface{}{
	