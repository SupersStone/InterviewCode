package erc20

import "github.com/ethereum/go-ethereum/crypto"

var (
	logTransferSig = []byte("Transfer(address,address,uint256)")
	logApprovalSig = []byte("Approval(address,address,uint256)")
	// LogTransferSigHash transfer func hash
	LogTransferSigHash = crypto.Keccak256Hash(logTransferSig)
	// LogApprovalSigHash approval func hash
	LogApprovalSigHash = crypto.Keccak256Hash(logApprovalSig)
)
