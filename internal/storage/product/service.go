package product

var Products []*Product

func Addproduct() bool {
	payload := []*Product{
		{
			Id:    1,
			Name:  "Cola-Cola",
			Price: 50000,
		},
		{
			Id:    2,
			Name:  "Zdrink",
			Price: 100000,
		},
		{
			Id:    3,
			Name:  "Sprit",
			Price: 80000,
		},
	}
	Products = append(Products, payload...)
	return true
}

func Showproduct() []*Product {
	return Products
}
