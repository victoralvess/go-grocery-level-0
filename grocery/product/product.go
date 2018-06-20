package product

type Product struct {
	Name	string	
	Price	float32
	InStock	int
}

func New(name string, price float32, inStock int) Product {
	p := Product { name, price, inStock }
	return p
}
