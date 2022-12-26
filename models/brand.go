package models

import "time"

type CreateBrandModel struct {
	Name         string `json:"name" binding:"required" minLength:"2" maxLength:"50" example:"Sherlock"`
	Country      string `json:"country" binding:"required" example:"USA, KOREA, GERMANY"`
	Manufacturer string `json:"manufacturer" binding:"required" example:"Berlin"`
	AboutBrand   string `json:"aboutbrand" binding:"required" example:"At Mercedes_Benz, our employes and communities are ate the heart of everything we do"`
}

type Brand struct {
	Id           string     `json:"brand_id"`
	Name         string     `json:"name" binding:"required" minLength:"2" maxLength:"50" example:"Sherlock"`
	Country      string     `json:"country" binding:"required" example:"USA, KOREA, GERMANY"`
	Manufacturer string     `json:"manufacturer" binding:"required" example:"Berlin"`
	AboutBrand   string     `json:"aboutbrand" binding:"required" example:"At Mercedes_Benz, our employes and communities are ate the heart of everything we do"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	DeletedAt    *time.Time `json:"-"`
}

type UpdateBrandModel struct {
	BrandId      string `json:"brand_id"`
	Name         string `json:"name" binding:"required" minLength:"2" maxLength:"50" example:"Sherlock"`
	Country      string `json:"country" binding:"required" example:"USA, KOREA, GERMANY"`
	Manufacturer string `json:"manufacturer" binding:"required" example:"Berlin"`
	AboutBrand   string `json:"aboutbrand" binding:"required" example:"At Mercedes_Benz, our employes and communities are ate the heart of everything we do"`
}