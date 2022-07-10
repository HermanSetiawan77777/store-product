package product_test

import (
	"testing"

	"github.com/hermansetiawan77777/technical-warungpintar/internal/storage/product"
)

func TestProductlist(t *testing.T) {
	productsadd := product.Addproduct()
	products := product.Showproduct()
	if !productsadd {
		t.Errorf("Failed generated Product")
		return
	}

	if products == nil {
		t.Errorf("Failed show List Product")
		return
	}
}

func TestAddProduct(t *testing.T) {
	productsadd := product.Addproduct()
	if !productsadd {
		t.Errorf("Failed generated Product")
		return
	}
}
