package orderHistory

type OrderHistory struct {
	ID        int    `json:"id"`
	Orderer   string `json:"orderer"`
	Product   int    `json:"product"`
	OrderedAt string `json:"ordered_at"`
}
