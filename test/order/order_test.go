package order_test

import (
	"testing"

	"github.com/hermansetiawan77777/technical-warungpintar/internal/storage/order"
)

func TestNewOrder(t *testing.T) {
	payload := &order.Bucket{
		Id:    1,
		Name:  "Cola-Cola",
		Price: 50000,
	}
	productsadd := order.Addorder("Hersti", payload)
	if payload.Id == 0 {
		t.Errorf("Please put item on you bucket")
		return
	}
	if productsadd == nil {
		t.Errorf("Failed create order !")
		return
	}
}
