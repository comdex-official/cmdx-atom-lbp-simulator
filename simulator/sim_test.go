package simulator_test

import (
	"testing"

	"github.com/public-awesome/lbp-simulator/simulator"
)

func TestSomething(t *testing.T) {
	sample := []byte(`{
		"weights": "90ucmdx,10uatom",
		"initial-deposit": "50000000000000ucmdx,125000000000uatom",
		"swap-fee": "0.003",
		"exit-fee": "0.001",
		"lbp-params": {
			"duration": "120h",
			"target-pool-weights": "1ucmdx,1uatom"
		}
	}`)
	poolMsg, err := simulator.CreatePoolMsg(sample)
	if err != nil {
		t.Fatal(err)
	}

	simulator.Simulate(poolMsg.PoolParams, poolMsg.PoolAssets, 100_000)
}
