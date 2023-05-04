package spec

// ListResponse is used to return list of products
type ListResponse struct {
	Products []Product `json:"products"`
}
