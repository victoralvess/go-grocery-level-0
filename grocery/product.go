package grocery

type product struct {
	Name	string	
	Price	float32
	InStock	int
}

func NewProduct(name string, price float32, inStock int) product {
	p := product { name, price, inStock }
	return p
}
