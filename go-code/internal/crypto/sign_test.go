package crypto

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func Test_signHash(t *testing.T) {
	// nonce 是随机的所以测试需要跳过
	if testing.Short() {
		t.Skip("short模式下会跳过该测试用例")
	}

	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test login sign",
			args: args{
				data: []byte(mockNonce("0xE11DeC63C6534C1AE5A89bfc740E44C477813D2F")),
			},
			want: "49f8a43bd4d653768be540a63af385a317516782eabc973115325d9d792230c2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SignHash(tt.args.data); !reflect.DeepEqual(hex.EncodeToString(got), tt.want) {
				t.Errorf("signHash() = %v, want %v", hex.EncodeToString(got), tt.want)
			}
		})
	}
}

func mockNonce(address string) string {
	uuid := uuid.New()
	val := uuid.String()

	return fmt.Sprintf(`Welcome to DAO-NFT!
	This request will not trigger a blockchain transaction or cost any gas fees.
	Your authentication status will reset after 24 hours.
	Wallet address:
	%s
	Nonce:
	%s}`, address, val)
}

func Test_sign(t *testing.T) {
	type args struct {
		priv     []byte
		signHash []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "",
			args: args{
				priv:     []byte{230, 242, 29, 42, 64, 132, 210, 180, 176, 150, 75, 235, 107, 27, 85, 80, 83, 23, 171, 159, 225, 246, 23, 152, 41, 229, 169, 205, 124, 53, 32, 194},
				signHash: []byte{19, 140, 92, 214, 59, 145, 126, 112, 248, 225, 150, 217, 35, 214, 0, 5, 150, 30, 121, 71, 107, 7, 148, 88, 161, 220, 40, 179, 141, 14, 158, 64},
			},
			want: []byte{175, 164, 195, 16, 100, 149, 237, 193, 16, 83, 49, 205, 46, 206, 49, 234, 200, 79, 200, 18, 191, 153, 0, 165, 189, 132, 221, 123, 7, 64, 111, 160, 125, 185, 230, 239, 52, 123, 222, 31, 188, 69, 117, 196, 200, 177, 201, 218, 238, 209, 171, 48, 147, 197, 18, 129, 186, 247, 28, 221, 62, 218, 4, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sign(tt.args.priv, tt.args.signHash); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sign() = %v, want %v", got, tt.want)
			}
		})
	}
}
