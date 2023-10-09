package services

import "github.com/dassudip2001/webapp/models"

type Catering struct {
	Name        string  `json:"name"`
	IsAvailable bool    `json:"is_available"`
	Price       float64 `json:"price"`
}

type CreateCateringRequest struct {
	Name        string  `json:"name"`
	IsAvailable bool    `json:"is_available"`
	Price       float64 `json:"price"`
}

func createResponseServices(servicesModel models.Catering) Catering {
	return Catering{
		Name:        servicesModel.Name,
		IsAvailable: servicesModel.IsAvailable,
		Price:       servicesModel.Price,
	}
}
