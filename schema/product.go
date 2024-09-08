package schema

type Product struct {
	Sku string `json:"sku"`
}

func (p *Product) GetSku() string {
	return p.Sku
}
