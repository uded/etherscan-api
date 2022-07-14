/*
 * Copyright (c) 2018 LI Zhennan
 * Copyright (c) 2022 Łukasz Rżanek
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_BlockReward(t *testing.T) {
	const (
		ethAns      = `{"blockNumber":"2165403","timeStamp":"1472533979","blockMiner":"0x13a06d3dfe21e0db5c016c03ea7d2509f7f8d1e3","blockReward":"5314181600000000000","uncles":[{"miner":"0xbcdfc35b86bedf72f0cda046a3c16829a2ef41d1","unclePosition":"0","blockreward":"3750000000000000000"},{"miner":"0x0d0c9855c722ff0c78f21e43aa275a5b8ea60dce","unclePosition":"1","blockreward":"3750000000000000000"}],"uncleInclusionReward":"312500000000000000"}`
		maticAns    = `{"blockNumber":"2165403","timeStamp":"1595322344","blockMiner":"0x0375b2fc7140977c9c76d45421564e354ed42277","blockReward":"0","uncles":[],"uncleInclusionReward":"0"}`
		bscAns      = `{"blockNumber":"2165403","timeStamp":"1605169045","blockMiner":"0x78f3adfc719c99674c072166708589033e2d9afe","blockReward":"0","uncles":[],"uncleInclusionReward":"0"}`
		avaxAns     = `{"blockNumber":"2165403","timeStamp":"1622557183","blockMiner":"0x0100000000000000000000000000000000000000","blockReward":"7963290000000000","uncles":[],"uncleInclusionReward":"0"}`
		ftmAns      = `{"blockNumber":"2165403","timeStamp":"1612983193","blockMiner":"0x0000000000000000000000000000000000000000","blockReward":"5225066000000000","uncles":[],"uncleInclusionReward":"0"}`
		cronosAns   = `{"blockNumber":"2165403","timeStamp":"1648843458","blockMiner":"0x4f87a3f99bd1e58d01de1c38b7f83cb967e816c2","blockReward":"49532402000000000000","uncles":[],"uncleInclusionReward":"0"}`
		arbitrumAns = `{"blockNumber":"2165403","timeStamp":"1634069963","blockMiner":"0x0000000000000000000000000000000000000000","blockReward":"2065103660121552","uncles":[],"uncleInclusionReward":"0"}`
	)

	type Data struct {
		BlockNumber int
		Ans         string
	}

	testData := map[string]Data{
		EthMainnet.CommonName:      {2165403, ethAns},
		MaticMainnet.CommonName:    {2165403, maticAns},
		BscMainnet.CommonName:      {2165403, bscAns},
		AvaxMainnet.CommonName:     {2165403, avaxAns},
		FantomMainnet.CommonName:   {2165403, ftmAns},
		CronosMainnet.CommonName:   {2165403, cronosAns},
		ArbitrumMainnet.CommonName: {2165403, arbitrumAns},
	}

	for _, network := range TestNetworks {
		if td, ok := testData[network.Network.CommonName]; ok {
			t.Run(network.Network.Name, func(t *testing.T) {
				reward, err := network.client.BlockReward(td.BlockNumber)
				assert.NoError(t, err)

				j, err := json.Marshal(reward)
				assert.NoError(t, err)
				assert.Equalf(t, td.Ans, string(j), "api.BlockReward not working, got %s, want %s", j, ethAns)
			})
		}
	}
}

func TestClient_BlockNumber(t *testing.T) {
	type Data struct {
		Timestamp int64
		AnsBefore int
		AnsAfter  int
	}

	testData := map[string]Data{
		// Note: All values taken from docs.etherscan.io/api-endpoints/blocks
		EthMainnet.CommonName:      {1578638524, 9251482, 9251483},
		MaticMainnet.CommonName:    {1601510400, 5164199, 5164200},
		BscMainnet.CommonName:      {1601510400, 946206, 946207},
		AvaxMainnet.CommonName:     {1609455600, 18960, 18961},
		FantomMainnet.CommonName:   {1609455600, 1632921, 1632921}, // it is the same!
		CronosMainnet.CommonName:   {1654034400, 3023556, 3023557},
		ArbitrumMainnet.CommonName: {1656626400, 16655953, 16655954},
	}

	for _, network := range TestNetworks {
		if td, ok := testData[network.Network.CommonName]; ok {
			t.Run(network.Network.Name, func(t *testing.T) {
				blockNumber, err := network.client.BlockNumber(td.Timestamp, "before")
				assert.NoError(t, err)
				assert.Equalf(t, td.AnsBefore, blockNumber, `api.BlockNumber(%d, "before") not working, got %d, want %d`, td.Timestamp, blockNumber, td.AnsBefore)

				blockNumber, err = network.client.BlockNumber(td.Timestamp, "after")
				assert.NoError(t, err)
				assert.Equalf(t, td.AnsAfter, blockNumber, `api.BlockNumber(%d, "after") not working, got %d, want %d`, td.Timestamp, blockNumber, td.AnsAfter)
			})
		}
	}
}
