// Package types provides core primitives for the operation
// of the Polaris protocol.
package types

import (
	"math/big"

	"github.com/polaris-project/go-polaris/common"
)

// Transaction is a data type representing a transfer of monetary value between addresses.
// A transactions does not necessarily imply the transfer of value between human peers, but also contracts.
type Transaction struct {
	AccountNonce uint64 `json:"nonce" gencodec:"required"` // Index in account transaction list

	Sender    *common.Address `json:"sender" gencodec:"required"`    // Transaction sender
	Recipient *common.Address `json:"recipient" gencodec:"required"` // Transaction recipient

	GasPrice *big.Int `json:"gas_price" gencodec:"required"` // Gas price in units equivalent to 0.000000001 of a single unit
	GasLimit uint64   `json:"gas_limit" gencodec:"required"` // Value of gas price willing to pay for transaction

	Payload []byte `json:"payload" gencodec:"required"` // Data sent with transaction (i.e. contract bytecode, message, etc...)

	Signature *Signature `json:"signature" gencodec:"required"` // ECDSA transaction signature

	Hash common.Hash `json:"hash" gencodec:"required"` // Transaction hash
}
