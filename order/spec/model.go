package spec

// Order ...
type Order struct {
	ID           int     `json:"id"`
	Items        []Item  `json:"items"`
	Amount       float64 `json:"amount"`
	Discount     float64 `json:"discount"`
	FinalAmount  float64 `json:"finalAmount"`
	Status       int     `json:"status"`
	OrderDate    string  `json:"orderDate"`
	DispatchDate string  `json:"dispatchDate,omitempty"`
}

// Item ...
type Item struct {
	ID       int `json:"id"`
	Quantity int `json:"quantity"`
}
