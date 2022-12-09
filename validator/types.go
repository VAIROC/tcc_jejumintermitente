package validator

// NetworkType ...
type NetworkType string

// List of NetworkType
const (
	Mainnet NetworkType = "Mainnet"
	Testnet NetworkType = "Testnet"
)

// AddressType ...
type AddressType string

// List of AddressType
const (
	Unknown            AddressType = "Unknown"
	Normal             AddressType = "N