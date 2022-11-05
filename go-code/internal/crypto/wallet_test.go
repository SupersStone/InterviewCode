package crypto

import (
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func Test_generatedWallet(t *testing.T) {
	if testing.Short() {
		t.Skip("short模式下会跳过该测试用例")
	}
	tests := []struct {
		name        string
		wantPriv    []byte
		wantAddress common.Address
	}{
		{
			name:        "test",
			wantPriv:    []byte{},
			wantAddress: [20]byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPriv, gotAddress := generatedWallet()
			if !reflect.DeepEqual(gotPriv, tt.wantPriv) {
				t.Errorf("generatedWallet() gotPriv = %v, want %v", gotPriv, tt.wantPriv)
			}
			if !reflect.DeepEqual(gotAddress, tt.wantAddress) {
				t.Errorf("generatedWallet() gotAddress = %v, want %v", gotAddress, tt.wantAddress)
			}
		})
	}
}
