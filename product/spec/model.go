package spec

// Product is the struct for product details
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	CategoryID  int     `json:"-"`
	Category    string  `json:"category"`
	AvlQuantity int     `json:"availableQuantity"`
}
