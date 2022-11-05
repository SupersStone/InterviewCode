package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"math/big"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

// NewSlice new uint64 slice from start to end and step in giving, the slice max len is count
func NewSlice[T int | int32 | int64 | uint | uint32 | uint64](start, count, end, step T) []T {
	res := []T{}
	for i := start; i <= end; i += step {
		if count == 0 {
			return res
		}

		res = append(res, start)
		start++
		count--
	}

	return res
}

// Timestamp to time.time
func TimestampToTime(timeData int64) string {
	timeTemplate := "2006-01-02 15:04:05"
	tm := time.Unix(timeData, 0)
	timeStr := tm.Format(timeTemplate)
	return timeStr
}

// Time.time to timestamp
func TimeToTimestamp(timeData string) int64 {
	var LOC, _ = time.LoadLocation("Asia/Shanghai")
	timeTmeplate := "2006-01-02 15:04:05"
	timData, _ := time.ParseInLocation(timeTmeplate, timeData, LOC)
	result := timData.Unix()
	return result
}

// Decimal to  Hexadecimal
func DecimalToHex(data string) string {

	// 字符串转int类型
	result, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		return ""
	}
	// int类型转16进制
	decimalData := strconv.FormatInt(result, 16)
	// 填充剩余的0
	zeroLen := 64 - len(decimalData)
	sum := "0x"
	for i := 0; i < zeroLen; i++ {
		sum += "0"
	}
	resultData := sum + string(decimalData)
	return resultData
}

// Hexadecimal to Decimal (Token id )
func HexTodecimal(hexData string) (int64, error) {
	// If the first two digits are 0x, then directly intercept the tone
	if hexData[:2] == "0x" {
		hexData = hexData[2:]
	}
	data, err := strconv.ParseUint(hexData, 16, 32)
	if err != nil {
		return 0, err
	}
	return int64(data), nil
}

// Address Length Checking
func AddressLenCheck(adddress string) bool {

	if adddress == "" {
		return false
	}

	if adddress[:2] != "0x" {
		return false
	}
	if len(adddress) != 42 {
		return false
	}
	return true

}

// Sign Checkin
func SignLenCheck(sign string) bool {
	if (sign)[:2] != "0x" {
		return false
	}

	if len(sign) != 132 {
		return false
	}
	return true
}

// IsValidAddress validate hex address
func IsValidAddress(iaddress interface{}) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := iaddress.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}

// IsZeroAddress validate if it's a 0 address
func IsZeroAddress(iaddress interface{}) bool {
	var address common.Address
	switch v := iaddress.(type) {
	case string:
		address = common.HexToAddress(v)
	case common.Address:
		address = v
	default:
		return false
	}

	zeroAddressBytes := common.FromHex("0x0000000000000000000000000000000000000000")
	addressBytes := address.Bytes()
	return reflect.DeepEqual(addressBytes, zeroAddressBytes)
}

// ToDecimal wei to decimals
func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result
}

// ToWei decimals to wei
func ToWei(iamount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}
