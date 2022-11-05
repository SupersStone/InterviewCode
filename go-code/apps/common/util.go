package common

import (
	"errors"
	"math/big"

	"github.com/shopspring/decimal"
)

// crypto info
const (
	ZeroAddress string = "0x0000000000000000000000000000000000000000"
	ZeroHash    string = "0x0000000000000000000000000000000000000000000000000000000000000000"
)

// HumanReadPriceFromStr big num to decimal
func HumanReadPriceFromStr(amount string, tokenDecimal int) (decimal.Decimal, error) {
	res := big.NewInt(0)
	amountBig, ok := res.SetString(amount, 10)
	if !ok {
		return decimal.Zero, errors.New("covert amount to bigInt err")
	}

	return decimal.NewFromBigInt(amountBig, 0).Div(decimal.New(10, int32(tokenDecimal))), nil
}
