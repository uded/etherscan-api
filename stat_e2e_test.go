/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_EtherTotalSupply(t *testing.T) {
	for _, network := range TestNetworks {
		t.Run(network.Network.CommonName, func(t *testing.T) {
			totalSupply, err := network.client.EtherTotalSupply()
			if network.Network.TokenName != "" {
				assert.NoError(t, err)

				if totalSupply.Int().Cmp(big.NewInt(100)) != 1 {
					t.Errorf("api.EtherTotalSupply not working, totalSupply is %s", totalSupply.Int().String())
				}
			} else {
				assert.Error(t, err)
			}

		})
	}
}

func TestClient_EtherLatestPrice(t *testing.T) {
	for _, network := range TestNetworks {
		t.Run(network.Network.CommonName, func(t *testing.T) {
			latest, err := network.client.EtherLatestPrice()
			assert.NoError(t, err)

			if latest.ETHBTC == 0 {
				t.Errorf("ETHBTC got 0")
			}
			if latest.ETHBTCTimestamp.Time().IsZero() {
				t.Errorf("ETHBTCTimestamp is zero")
			}
			if latest.ETHUSD == 0 {
				t.Errorf("ETHUSD got 0")
			}
			if latest.ETHUSDTimestamp.Time().IsZero() {
				t.Errorf("ETHUSDTimestamp is zero")
			}
		})
	}
}

func TestClient_TokenTotalSupply(t *testing.T) {
	totalSupply, err := api.TokenTotalSupply("0x57d90b64a1a57749b0f932f1a3395792e12e7055")
	assert.NoError(t, err)

	if totalSupply.Int().Cmp(big.NewInt(100)) != 1 {
		t.Errorf("api.TokenTotalSupply not working, totalSupply is %s", totalSupply.Int().String())
	}
}
