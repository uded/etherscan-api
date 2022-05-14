/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan

const (
	// // Ethereum public networks

	// EthMainnet Ethereum mainnet for production
	EthMainnet Network = "https://api.etherscan.io/api?"
	// EthRopsten Testnet(POW)
	EthRopsten Network = "https://api-ropsten.etherscan.io/api?"
	// EthKovan Testnet(POA)
	EthKovan Network = "https://api-kovan.etherscan.io/api?"
	// EthRinkby Testnet(CLIQUE)
	EthRinkby Network = "https://api-rinkeby.etherscan.io/api?"
	// EthGoerli Testnet(CLIQUE)
	EthGoerli Network = "https://api-goerli.etherscan.io/api?"
	// EthTobalaba Testnet
	EthTobalaba Network = "https://api-tobalaba.etherscan.io/api?"
	// MaticMainnet Matic mainnet for production
	MaticMainnet Network = "https://api.polygonscan.com/api?"
	// MaticTestnet Matic testnet for development
	MaticTestnet Network = "https://api-testnet.polygonscan.com/api?"
	// BscMainnet Bsc mainnet for production
	BscMainnet Network = "https://api.bscscan.com/api?"
	// BscTestnet Bsc testnet for development
	BscTestnet Network = "https://api-testnet.bscscan.com/api?"
)

// Network is ethereum network type (mainnet, ropsten, etc)
type Network string

// SubDomain returns the subdomain of  etherscan API
// via n provided.
func (n Network) Domain() (sub string) {
	return string(n)
}
