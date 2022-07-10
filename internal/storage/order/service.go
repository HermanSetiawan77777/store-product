package order

var Orders []*Order

func Addorder(username string, bucketitem *Bucket) *Order {
	payload1 := &Order{
		Idtrans:       len(Orders) + 1,
		Username:      username,
		Productbucket: *bucketitem,
	}
	Orders = append(Orders, payload1)

	return payload1
}

func Showorder(id string) []*Order {
	var Orderuser []*Order
	for _, v := range Orders {
		if v.Username == id {
			Orderuser = append(Orderuser, v)
		}
	}
	return Orderuser
}
