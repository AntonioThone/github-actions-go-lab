package main

import (
	"encoding/json"
	"net/http"

	"github-actions-go-lab/internal/conta"
)

type Request struct {
	Valor        float64 `json:"valor"`
	TaxaServico  float64 `json:"taxaServico"`
	Pessoas      int     `json:"pessoas"`
}

type Response struct {
	ValorPorPessoa float64 `json:"valorPorPessoa"`
}

func calcular(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req Request

	json.NewDecoder(r.Body).Decode(&req)

	resultado := conta.CalcularConta(
		req.Valor,
		req.TaxaServico,
		req.Pessoas,
	)

	resp := Response{
		ValorPorPessoa: resultado,
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(resp)
}

func main() {

	http.HandleFunc(
		"/calcular",
		calcular,
	)

	http.ListenAndServe(
		":8080",
		nil,
	)
}