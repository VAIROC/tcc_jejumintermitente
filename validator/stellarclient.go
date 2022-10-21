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
	url := fmt.Sprintf("%s/accounts/%s", c.Endpoint, addr)
	resp, err := httpclient.Get(url)
	if err != nil {
		return Unknown, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return Unknown, nil