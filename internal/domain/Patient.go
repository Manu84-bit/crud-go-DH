package domain

import "time"

type Patient struct {
	Id            int    `json:"id"`
	Name          string `json:"name" binding:"required"`
	LastName      int    `json:"last_name" binding:"required"`
	Address       string `json:"address" binding:"required"`
	DNI           int    `json:"dni" binding:"required"`
	DischargeDate time.Time `json:"discharge_date" binding:"required"`
}