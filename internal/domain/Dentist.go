package domain

type Dentist struct {
	Id          int    `json:"id"`
	Name        string `json:"name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	LicenseCode string `json:"license" binding:"required"`
}