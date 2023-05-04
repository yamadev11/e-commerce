package spec

// UpdateRequest is used to update to product quantity.
type UpdateRequest struct {
	ID       int `json:"id"`
	Quantity int `json:"quantity"`
}
