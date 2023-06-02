/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

const apiKeyEnvName = "ETHERSCAN_API_KEY"

var (
	// api test client for many test cases
	api *Client
	// bucket default rate limiter
	bucket *Bucket
	// apiKey etherscan API key
	apiKey string

	TestNetworks []Networks
)

type Networks struct {
	Network Network
	APIKey  string
	client  *Client
}

func init() {
	apiKey = os.Getenv(apiKeyEnvName)
	networks := os.Getenv("NETWORKS")
	if apiKey == "" {
		panic(fmt.Sprintf("API key is empty, set env variable %q with a valid API key to proceed.", apiKeyEnvName))
	}
	bucket = NewBucket(500 * time.Millisecond)

	networksParsed := strings.Split(networks, ",") // parse all values as slice
	for _, networkParsed := range networksParsed {
		typeAndKey := strings.Split(networkParsed, ":") // split the network type and api key
		if len(typeAndKey) != 2 {
			// return errors.New("invalid network format - should be <network>:<api-key>")
		}
		network, err := ParseNetworkName(strings.TrimSpace(typeAndKey[0]))
		if err != nil {
			// return err
		}

		api = New(network, typeAndKey[1])
		api.Verbose = true
		api.BeforeRequest = func(module string, action string, param map[string]interface{}) error {
			bucket.Take()
			return nil
		}

		TestNetworks = append(TestNetworks, Networks{network, typeAndKey[1], api})
	}

	api = New(EthMainnet, apiKey)
	api.Verbose = true
	api.BeforeRequest = func(module string, action string, param map[string]interface{}) error {
		bucket.Take()
		return nil
	}
}

// Bucket is a simple and easy rate limiter
// Use NewBucket() to construct one.
type Bucket struct {
	bucket     chan bool
	refillTime time.Duration
}

// NewBucket factory of Bucket
func NewBucket(refillTime time.Duration) (b *Bucket) {
	b = &Bucket{
		bucket:     make(chan bool),
		refillTime: refillTime,
	}

	go b.fillRoutine()

	return
}

// Take a action token from bucket,
// blocks if there is currently no token left.
func (b *Bucket) Take() {
	<-b.bucket
}

// fill a action token into bucket,
// no-op if the bucket is currently full
func (b *Bucket) fill() {
	b.bucket <- true
}

func (b *Bucket) fillRoutine() {
	ticker := time.NewTicker(b.refillTime)

	for range ticker.C {
		b.fill()
	}
}

// noError checks for testing error
func noError(t *testing.T, err error, msg string) {
	if err != nil {
		t.Fatalf("%s: %v", msg, err)
	}
}
