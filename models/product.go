package models

// Product refers to the product model on database
type Product struct {
	_id         string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Price       int    `json:"price,omitempty"`
}
