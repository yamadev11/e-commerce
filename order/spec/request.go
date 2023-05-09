package spec

// GetRequest is used to get details of provided ID
type GetRequest struct {
	ID int `json:"id"`
}

// UpdateRequest is used to update the order status
type UpdateRequest struct {
	ID     int `json:"id"`
	Status int `json:"status"`
}

// CreateRequest is used to create new order
type CreateRequest struct {
	Items []Item `json:"items"`
}
