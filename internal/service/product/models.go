package product

var allProducts = []Product{
	{Title: "Product - 1"},
	{Title: "Product - 2"},
	{Title: "Product - 3"},
	{Title: "Product - 4"},
	{Title: "Product - 5"},
}

type Product struct {
	Title string
}

func (p *Product) String() string {
	return p.Title
}
