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
		{EthMainnet.Name, EthMainnet, false},
		{EthMainnet.CommonName, EthMainnet, false},
		{EthRopsten.Name, EthRopsten, false},
		{EthRopsten.CommonName, EthRopsten, false},
		{EthKovan.Name, EthKovan, false},
		{EthKovan.CommonName, EthKovan, false},
		{EthRinkby.Name, EthRinkby, false},
		{EthRinkby.CommonName, EthRinkby, false},
		{EthGoerli.Name, EthGoerli, false},
		{EthGoerli.CommonName, EthGoerli, false},
		{EthTobalaba.Name, EthTobalaba, false},
		{EthTobalaba.CommonName, EthTobalaba, false},
		{MaticMainnet.Name, MaticMainnet, false},
		{MaticMainnet.CommonName, MaticMainnet, false},
		{MaticTestnet.Name, MaticTestnet, false},
		{MaticTestnet.CommonName, MaticTestnet, false},
		{BscMainnet.Name, BscMainnet, false},
		{BscMainnet.CommonName, BscMainnet, false},
		{BscTestnet.Name, BscTestnet, false},
		{BscTestnet.CommonName, BscTestnet, false},
		{AvaxMainnet.Name, AvaxMainnet, false},
		{AvaxMainnet.CommonName, AvaxMainnet, false},
		{AvaxTestnet.Name, AvaxTestnet, false},
		{AvaxTestnet.CommonName, AvaxTestnet, false},
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
