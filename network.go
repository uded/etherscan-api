/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan

var (
	// EthMainnet Ethereum mainnet for production
	EthMainnet Network = Network{"Ethereum", "main", "https://api.etherscan.io/api?"}
	// EthRopsten Testnet(POW)
	EthRopsten Network = Network{"Ethereum Ropsten", "test", "https://api-ropsten.etherscan.io/api?"}
	// EthKovan Testnet(POA)
	EthKovan Network = Network{"Ethereum Kovan", "test", "https://api-kovan.etherscan.io/api?"}
	// EthRinkby Testnet(CLIQUE)
	EthRinkby Network = Network{"Ethereum Rinkby", "test", "https://api-rinkeby.etherscan.io/api?"}
	// EthGoerli Testnet(CLIQUE)
	EthGoerli Network = Network{"Ethereum Goerli", "test", "https://api-goerli.etherscan.io/api?"}
	// EthTobalaba Testnet
	EthTobalaba Network = Network{"Ethereum Tobalaba", "test", "https://api-tobalaba.etherscan.io/api?"}
	// MaticMainnet Matic mainnet for production
	MaticMainnet Network = Network{"Polygon", "main", "https://api.polygonscan.com/api?"}
	// MaticTestnet Matic testnet for development
	MaticTestnet Network = Network{"Polygon Mumbai", "test", "https://api-testnet.polygonscan.com/api?"}
	// BscMainnet Bsc mainnet for production
	BscMainnet Network = Network{"Binance", "main", "https://api.bscscan.com/api?"}
	// BscTestnet Bsc testnet for development
	BscTestnet Network = Network{"Binance test", "test", "https://api-testnet.bscscan.com/api?"}
)

// Network is ethereum network type (mainnet, ropsten, etc)
type Network struct {
	Name    string // Name of the network or chain
	Type    string // Type of the network, either main or test
	baseURL string // baseURL for the API client
}

// Domain returns the subdomain of etherscan API via n provided.
func (n Network) Domain() (sub string) {
	return n.baseURL
}
