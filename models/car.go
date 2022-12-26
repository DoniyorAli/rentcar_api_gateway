package models

import "time"

type Car struct {
	CarId     string     `json:"car_id"`
	Model     string     `json:"model" binding:"required" example:"Tesla Company"`
	Color     string     `json:"color" example:"Dark-Black"`
	CarType   string     `json:"car_type"   example:"electro_car"`
	Mileage   string     `json:"mileage"   example:"360-km"`
	Year      string     `json:"year" example:"2020"`
	Price     string     `json:"price" example:"70$"`
	BrandId   string     `json:"brand_id" example:"1"`
	CreatedAt time.Time  `json:"created_at"`
	UpdateAt  *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

type CreateCarModel struct {
	Model   string `json:"model" binding:"required" example:"Tesla Company"`
	Color   string `json:"color" example:"Dark-Black"`
	CarType string `json:"car_type"   example:"electro_car"`
	Mileage string `json:"mileage"   example:"360-km"`
	Year    string `json:"year" example:"2020"`
	Price   string `json:"price" example:"70$"`
	BrandId string `json:"brand_id" example:"1"`
}

type PackedCarModel struct {
	CarId     string     `json:"car_id"`
	Model     string     `json:"fullname" binding:"required" example:"Tesla"`
	Color     string     `json:"color" example:"White"`
	CarType   string     `json:"car_type"   example:"electro_car"`
	Mileage   string     `json:"mileage"   example:"540-km"`
	Year      string     `json:"year" example:"2022"`
	Price     string     `json:"price" example:"80$"`
	Brand     Brand      `json:"brand_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdateAt  *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

type UpdateCarModel struct {
	CarId     string     `json:"car_id"`
	Model     string     `json:"model" binding:"required" example:"Tesla Company"`
	Color     string     `json:"color" example:"Dark-Black"`
	CarType   string     `json:"car_type"   example:"electro_car"`
	Mileage   string     `json:"mileage"   example:"360-km"`
	Year      string     `json:"year" example:"2020"`
	Price     string     `json:"price" example:"70$"`
	BrandId   string     `json:"brand_id" example:"1"`
}
