package order

type Order struct {
	Idtrans       int
	Username      string
	Productbucket Bucket
}

type Bucket struct {
	Id    int
	Name  string
	Price int
}
