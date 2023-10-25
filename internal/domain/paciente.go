package domain

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name" binding:"required"`
	Quantity    int     `json:"quantity" binding:"required"`
	CodeValue   string  `json:"code" binding:"required"`
	IsPublished bool    `json:"is_published" binding:"required"`
	Expiration  string  `json:"expiration_date" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}