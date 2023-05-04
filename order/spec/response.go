package spec

// GetOrderResponse holds the response for GetOrder API
type GetOrderResponse struct {
	ID           int         `json:"id"`
	Items        []OrderItem `json:"items"`
	Amount       float64     `json:"amount"`
	Discount     float64     `json:"discount"`
	FinalAmount  float64     `json:"finalAmount"`
	Status       string      `json:"status"`
	OrderDate    string      `json:"orderDate"`
	DispatchDate string      `json:"dispatchDate,omitempty"`
}

// OrderItem contains the details of the ordered item
type OrderItem struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
	Quantity int     `json:"quantity"`
}
