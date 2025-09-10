package models

type Product struct {
	Name     string `json:"name"`
	SKU      string `json:"sku"`
	Quantity int    `json:"quantity"`
	Location string `json:"location"`
	Status   string `json:"status"`
}
