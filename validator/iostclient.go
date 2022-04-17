// endpoint = chain.config["endpoint"]
// url = "#{endpoint}/getAccount/#{account}/false"

package validator

import (
	"fmt"

	"github.com/LanfordCai/ava/httpclient"
)

// IOSTClient ...
type IOSTClient struct {
	Endpoint string
}

// GetAccount ...
func (c *IOSTClient) GetAccount(addr string) (AddressType, error) {
	url := fmt.Sprintf("%s/getAccount/%s/false"