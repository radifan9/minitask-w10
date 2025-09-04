package models

// Request body for PATCH
type UpdatePriceRequest struct {
	Id    int `json:"id,omitempty"`
	Price int `json:"price" binding:"required"`
}
