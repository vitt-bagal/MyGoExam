package product

type productType int16

const (
	Grocery productType = iota + 1
	Electronics
	Kitchen
)

type Product struct {
	Name   string
	Type   productType
	Price  float64
	Seller string
}

func New() Product {
	prod := Product{"Tv", 2, 15550, "ABC"}
	return prod
}

func (p *Product) ApplyDiscount(per float64) {
	discountedPrice := (p.Price * per) / 100
	p.Price = p.Price - discountedPrice
}
