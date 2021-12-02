package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/public-awesome/lbp-simulator/simulator"
)

func main() {
	fmt.Println("listening: 8080")
	http.HandleFunc("/api/simulate", func(rw http.ResponseWriter, r *http.Request) {
		req := SimulateRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		bz, _ := req.Marshal()
		poolMsg, err := simulator.CreatePoolMsg(bz)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		rw.Header().Add("content-type", "application/json")
		resp, err := simulator.Simulate(poolMsg.PoolParams, poolMsg.PoolAssets, req.Volume)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(rw).Encode(resp)
	})
	http.ListenAndServe(":8080", nil)
}

type SimulateRequest struct {
	Duration      string `json:"duration"`
	InitialWeight struct {
		CMDX int `json:"cmdx"`
		Atom  int `json:"atom"`
	} `json:"initialWeight"`
	EndWeight struct {
		CMDX int `json:"cmdx"`
		Atom  int `json:"atom"`
	} `json:"endWeight"`
	Volume  int64 `json:"volume"`
	Deposit struct {
		CMDX int `json:"cmdx"`
		Atom  int `json:"atom"`
	} `json:"deposit"`
	Fees struct {
		Swap string `json:"swap"`
		Exit string `json:"exit"`
	} `json:"fees"`
}

func (r SimulateRequest) Marshal() ([]byte, error) {
	tpl := `{
		"weights": "%sucmdx,%suatom",
		"initial-deposit": "%sucmdx,%suatom",
		"swap-fee": "%s",
		"exit-fee": "%s",
		"lbp-params": {
			"duration": "%s",
			"target-pool-weights": "%sucmdx,%suatom"
		}
	}
	`
	tmp := fmt.Sprintf(tpl,
		strconv.Itoa(r.InitialWeight.CMDX), strconv.Itoa(r.InitialWeight.Atom),
		strconv.Itoa(r.Deposit.CMDX*1_000_000), strconv.Itoa(r.Deposit.Atom*1_000_000),
		r.Fees.Swap, r.Fees.Exit,
		r.Duration,
		strconv.Itoa(r.EndWeight.CMDX), strconv.Itoa(r.EndWeight.Atom),
	)
	return []byte(tmp), nil
}
