package crypto

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

// SignHash hash sign data
func SignHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), string(data))
	sighash := crypto.Keccak256([]byte(msg))

	return sighash
}

func sign(priv, signHash []byte) []byte {
	privateKey, err := crypto.ToECDSA(priv)
	if err != nil {
		log.Fatal(err)
	}

	signature, err := crypto.Sign(signHash, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	return signature
}
