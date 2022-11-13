package product

type Product struct {
	ID               int             `json:"id"`
	Name             string          `json:"name"`
	Price            int             `json:"price"`
	MyOrderHistories []*OrderHistory `json:"my_order_histories"`
}

type OrderHistory struct {
	ID        int    `json:"id"`
	OrderedAt string `json:"ordered_at"`
}
