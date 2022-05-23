package etherscan

import (
	"reflect"
	"testing"
)

func TestParseNetworkName(t *testing.T) {
	tests := []struct {
		name    string
		want    Network
		wantErr bool
	}{
		{"ethmainnet", EthMainnet, false},
		{"eth main net", Network{}, true},
		{"ethereum", EthMainnet, false},
		{"eth", EthMainnet, false},
		{"ethropsten", EthRopsten, false},
		{"ropsten", EthRopsten, false},
		{"ethkovan", EthKovan, false},
		{"ethrinkby", EthRinkby, false},
		{"ethgoerli", EthGoerli, false},
		{"ethtobalaba", EthTobalaba, false},
		{"maticmainnet", MaticMainnet, false},
		{"polygon", MaticMainnet, false},
		{"matic", MaticMainnet, false},
		{"matictestnet", MaticTestnet, false},
		{"mumbai", MaticTestnet, false},
		{"bscmainnet", BscMainnet, false},
		{"binance", BscMainnet, false},
		{"bsctestnet", BscTestnet, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseNetworkName(tt.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseNetworkName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseNetworkName() got = %v, want %v", got, tt.want)
			}
		})
	}
}
