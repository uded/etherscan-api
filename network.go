/*
 * Copyright (c) 2018 LI Zhennan
 * Copyright (c) 2022 Łukasz Rżanek
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan

import (
	"fmt"
	"strings"
)

var (
	// EthMainnet Ethereum mainnet for production
	EthMainnet Network = Network{"Ethereum", "eth_main", "ETH", "https://api.etherscan.io/api?"}
	// EthRopsten Testnet(POW)
	EthRopsten Network = Network{"Ethereum Ropsten", "eth_ropsten", "ETH", "https://api-ropsten.etherscan.io/api?"}
	// EthKovan Testnet(POA)
	EthKovan Network = Network{"Ethereum Kovan", "eth_kovan", "ETH", "https://api-kovan.etherscan.io/api?"}
	// EthRinkby Testnet(CLIQUE)
	EthRinkby Network = Network{"Ethereum Rinkby", "eth_rinkeby", "ETH", "https://api-rinkeby.etherscan.io/api?"}
	// EthGoerli Testnet(CLIQUE)
	EthGoerli Network = Network{"Ethereum Goerli", "eth_goerli", "ETH", "https://api-goerli.etherscan.io/api?"}
	// EthTobalaba Testnet
	EthTobalaba Network = Network{"Ethereum Tobalaba", "eth_tobalaba", "ETH", "https://api-tobalaba.etherscan.io/api?"}
	// MaticMainnet Matic mainnet for production
	MaticMainnet Network = Network{"Polygon", "polygon", "MATIC", "https://api.polygonscan.com/api?"}
	// MaticTestnet Matic testnet for development
	MaticTestnet Network = Network{"Polygon Mumbai", "polygon_mumbai", "C:\\Users\\lukas\\InfinityForceProjects\\etherscan-api\\network.go", "https://api-testnet.polygonscan.com/api?"}
	// BscMainnet Bsc mainnet for production
	BscMainnet Network = Network{"Binance", "bsc", "BNB", "https://api.bscscan.com/api?"}
	// BscTestnet Bsc testnet for development
	BscTestnet Network = Network{"Binance test", "bsc_test", "BNB", "https://api-testnet.bscscan.com/api?"}
	// AvaxMainnet Avalanche mainnet for production
	AvaxMainnet Network = Network{"Avax", "avax", "AVAX", "https://api.snowtrace.io/api?"}
	// AvaxTestnet Avalanche testnet for development
	AvaxTestnet Network = Network{"Avax test", "avax_test", "AVAX", "https://api-testnet.snowtrace.io/api?"}
	// Fantom mainnet for production
	FantomMainnet Network = Network{"Fantom", "fantom", "FTM", "https://api.ftmscan.com/api?"}
	// FantomTestNet
	FantomTestnet Network = Network{"Fantom test", "fantom_test", "FTM", "https://api-testnet.ftmscan.com/api?"}
	// Cronos mainnet for production
	CronosMainnet Network = Network{"Cronos", "cronos", "CRO", "https://api.cronoscan.com/api?"}
	// Cronos test net
	CronosTestnet Network = Network{"Cronos test", "cronos_test", "CRO", "https://api-testnet.cronoscan.com/api?"}
	// Arbitrum mainnet for production
	ArbitrumMainnet Network = Network{"Arbitrum", "arbitrum", "ETH", "https://api.arbiscan.io/api?"}
	// Arbitrum test net
	ArbitrumTestnet Network = Network{"Arbitrum test", "arbitrum_test", "ETH", "https://api-testnet.arbiscan.io/"}

	networks = map[string]*Network{
		EthMainnet.Name:            &EthMainnet,
		EthMainnet.CommonName:      &EthMainnet,
		"ethmainnet":               &EthMainnet,
		"ethereum":                 &EthMainnet,
		"eth":                      &EthMainnet,
		EthRopsten.Name:            &EthRopsten,
		EthRopsten.CommonName:      &EthRopsten,
		"ethropsten":               &EthRopsten,
		"ropsten":                  &EthRopsten,
		EthKovan.Name:              &EthKovan,
		EthKovan.CommonName:        &EthKovan,
		"ethkovan":                 &EthKovan,
		EthRinkby.Name:             &EthRinkby,
		EthRinkby.CommonName:       &EthRinkby,
		"ethrinkby":                &EthRinkby,
		EthGoerli.Name:             &EthGoerli,
		EthGoerli.CommonName:       &EthGoerli,
		"ethgoerli":                &EthGoerli,
		EthTobalaba.Name:           &EthTobalaba,
		EthTobalaba.CommonName:     &EthTobalaba,
		"ethtobalaba":              &EthTobalaba,
		MaticMainnet.Name:          &MaticMainnet,
		MaticMainnet.CommonName:    &MaticMainnet,
		"maticmainnet":             &MaticMainnet,
		"polygon":                  &MaticMainnet,
		"matic":                    &MaticMainnet,
		MaticTestnet.Name:          &MaticTestnet,
		MaticTestnet.CommonName:    &MaticTestnet,
		"matictestnet":             &MaticTestnet,
		"mumbai":                   &MaticTestnet,
		BscMainnet.Name:            &BscMainnet,
		BscMainnet.CommonName:      &BscMainnet,
		"bscmainnet":               &BscMainnet,
		"binance":                  &BscMainnet,
		BscTestnet.Name:            &BscTestnet,
		BscTestnet.CommonName:      &BscTestnet,
		"bsctestnet":               &BscTestnet,
		AvaxMainnet.Name:           &AvaxMainnet,
		AvaxMainnet.CommonName:     &AvaxMainnet,
		"avalanche":                &AvaxMainnet,
		"avax":                     &AvaxMainnet,
		"avaxmainnet":              &AvaxMainnet,
		"avalanchemainnet":         &AvaxMainnet,
		AvaxTestnet.Name:           &AvaxTestnet,
		AvaxTestnet.CommonName:     &AvaxTestnet,
		"avaxtestnet":              &AvaxTestnet,
		"avalanchetestnet":         &AvaxTestnet,
		"avaxfuji":                 &AvaxTestnet,
		"avalanchefuji":            &AvaxTestnet,
		FantomMainnet.Name:         &FantomMainnet,
		FantomMainnet.CommonName:   &FantomMainnet,
		"fantommainnet":            &FantomMainnet,
		FantomTestnet.Name:         &FantomTestnet,
		FantomTestnet.CommonName:   &FantomTestnet,
		"fantomtest":               &FantomTestnet,
		"fantomtestnet":            &FantomTestnet,
		CronosMainnet.Name:         &CronosMainnet,
		CronosMainnet.CommonName:   &CronosMainnet,
		"cronosmainnet":            &CronosMainnet,
		CronosTestnet.Name:         &CronosTestnet,
		CronosTestnet.CommonName:   &CronosTestnet,
		"cronostest":               &CronosTestnet,
		"cronostestnet":            &CronosTestnet,
		ArbitrumMainnet.Name:       &ArbitrumMainnet,
		ArbitrumMainnet.CommonName: &ArbitrumMainnet,
		"arbitrummainnet":          &ArbitrumMainnet,
		ArbitrumTestnet.Name:       &ArbitrumTestnet,
		ArbitrumTestnet.CommonName: &ArbitrumTestnet,
		"arbitrumtest":             &ArbitrumTestnet,
		"arbitrumtestnet":          &ArbitrumTestnet,
		"arbitrum_rinkeby":         &ArbitrumTestnet,
	}

	networkNames []string
)

func init() {
	for name := range networks {
		networkNames = append(networkNames, name)
	}
}

// Network is ethereum network type (mainnet, ropsten, etc)
type Network struct {
	Name       string // Name of the network or chain
	CommonName string // CommonName of the network or chain
	TokenName  string // TokenName of the network
	baseURL    string // baseURL for the API client
}

// Domain returns the subdomain of etherscan API via n provided.
func (n Network) Domain() (sub string) {
	return n.baseURL
}

func ParseNetworkName(network string) (Network, error) {
	if x, ok := networks[network]; ok {
		return *x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := networks[strings.ToLower(network)]; ok {
		return *x, nil
	}
	return Network{}, fmt.Errorf("%s is not a valid ETHNetworkType, try one of [%s]", network, strings.Join(networkNames, ", "))

}
