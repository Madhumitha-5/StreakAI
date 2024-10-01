package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Params struct {
	Array  []int
	Target int
}

type Response struct {
	Solution [][]int
}

func FindPairs(params *Params) *Response {
	resp := [][]int{}
	for i := range params.Array {
		for j := i + 1; j < len(params.Array); j++ {
			if params.Array[i]+params.Array[j] == params.Target {
				out := []int{i, j}
				resp = append(resp, out)
			}
		}
	}
	return &Response{Solution: resp}
}

func findPairsHandler(w http.ResponseWriter, r *http.Request) {
	var params Params
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}
	response := FindPairs(&params)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/find-pairs", findPairsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
