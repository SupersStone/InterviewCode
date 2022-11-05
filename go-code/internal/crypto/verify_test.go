package crypto

import (
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func Test_verify(t *testing.T) {
	type args struct {
		from     string
		sigHex   string
		signHash []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test login sign",
			args: args{
				from:     "0xE11DeC63C6534C1AE5A89bfc740E44C477813D2F",
				sigHex:   signHex(),
				signHash: []byte{19, 140, 92, 214, 59, 145, 126, 112, 248, 225, 150, 217, 35, 214, 0, 5, 150, 30, 121, 71, 107, 7, 148, 88, 161, 220, 40, 179, 141, 14, 158, 64},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.sigHex)
			if got := verify(tt.args.from, tt.args.sigHex, tt.args.signHash); got != tt.want {
				t.Errorf("verify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func signHex() string {
	bytes := []byte{175, 164, 195, 16, 100, 149, 237, 193, 16, 83, 49, 205, 46, 206, 49, 234, 200, 79, 200, 18, 191, 153, 0, 165, 189, 132, 221, 123, 7, 64, 111, 160, 125, 185, 230, 239, 52, 123, 222, 31, 188, 69, 117, 196, 200, 177, 201, 218, 238, 209, 171, 48, 147, 197, 18, 129, 186, 247, 28, 221, 62, 218, 4, 0, 1}
	return hexutil.Encode(bytes)
}

func Test_verifySign(t *testing.T) {
	type args struct {
		from      string
		signature []byte
		hash      []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test login sign",
			args: args{
				from:      "0xE11DeC63C6534C1AE5A89bfc740E44C477813D2F",
				signature: []byte{175, 164, 195, 16, 100, 149, 237, 193, 16, 83, 49, 205, 46, 206, 49, 234, 200, 79, 200, 18, 191, 153, 0, 165, 189, 132, 221, 123, 7, 64, 111, 160, 125, 185, 230, 239, 52, 123, 222, 31, 188, 69, 117, 196, 200, 177, 201, 218, 238, 209, 171, 48, 147, 197, 18, 129, 186, 247, 28, 221, 62, 218, 4, 0, 1},
				hash:      []byte{19, 140, 92, 214, 59, 145, 126, 112, 248, 225, 150, 217, 35, 214, 0, 5, 150, 30, 121, 71, 107, 7, 148, 88, 161, 220, 40, 179, 141, 14, 158, 64},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VerifySign(tt.args.from, tt.args.signature, tt.args.hash); got != tt.want {
				t.Errorf("verifySign() = %v, want %v", got, tt.want)
			}
		})
	}
}
