package models

type Cotacao struct {
	BRLAmount float64 `json:"brl_amount"` // Valor em reais (R$)
	USDAmount float64 `json:"usd_amount"` // Valor em dólares (USD)
}
