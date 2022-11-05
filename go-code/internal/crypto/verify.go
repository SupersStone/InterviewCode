package crypto

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

// enum
const (
	Evm = iota + 1
	Solana
)

// Verify a interface for verify func.
type Verify interface {
	Verify(from string, signedData string)
}

// Verifier a impl for Verify
type Verifier struct {
	chainID int
}

// NewEVMVerifier new Verifier
func NewEVMVerifier() *Verifier {
	return &Verifier{
		chainID: Evm,
	}
}

// NewVerifier new Verifier
func NewVerifier(chainType int) *Verifier {
	return &Verifier{
		chainID: chainType,
	}
}

// Verify verify sign data's siger equals from?
func (v *Verifier) Verify(from, sigHex string, signHash []byte) bool {
	switch v.chainID {
	case Evm:
		return verify(from, sigHex, signHash)
	default:
		return false
	}
}

// VerifySignature verify sign data's siger equals from?
func (v *Verifier) VerifySignature(from string, signature, signHash []byte) bool {
	switch v.chainID {
	case Evm:
		return VerifySign(from, signature, signHash)
	default:
		return false
	}
}

// Verify verify signature
func verify(from, sigHex string, hash []byte) bool {
	fromAddr := common.HexToAddress(from)

	sig := hexutil.MustDecode(sigHex)
	if sig[64] != 27 && sig[64] != 28 {
		return false
	}

	sig[64] -= 27

	pubKey, err := crypto.SigToPub(hash, sig)
	if err != nil {
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	return strings.EqualFold(fromAddr.Hex(), recoveredAddr.Hex())
}

// VerifySign verify signature
func VerifySign(from string, signature, hash []byte) bool {
	fromAddr := common.HexToAddress(from)

	pubKey, err := crypto.SigToPub(hash, signature)
	if err != nil {
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	return strings.EqualFold(fromAddr.Hex(), recoveredAddr.Hex())
}

// StructData EIP712结构化签名.
func StructData(typedData apitypes.TypedData) ([]byte, error) {
	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return nil, fmt.Errorf("domain hash fail, %w", err)
	}

	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return nil, fmt.Errorf("message hash fail, %w", err)
	}

	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	sighash := crypto.Keccak256(rawData)

	return sighash, nil
}
