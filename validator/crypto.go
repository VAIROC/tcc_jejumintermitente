
package validator

import (
	"github.com/btcsuite/btcutil/bech32"
)

func bech32AddrDecode(address string) (string, []byte, error) {
	hrp, data, err := bech32.Decode(address)
	if err != nil || len(data) < 1 {
		return "", nil, err
	}
	program, err := bech32.ConvertBits(data, 5, 8, false)
	if err != nil {
		return "", nil, err
	}
	return hrp, program, nil
}

func segwitAddrDecode(address string) (string, byte, []byte, error) {
	hrp, data, err := bech32.Decode(address)
	if err != nil || len(data) < 1 {
		return "", 255, nil, err