package controllers

import services "song-library/internal/services"

type Controller struct {
	Service services.ServiceInterface
}

func NewController(service services.ServiceInterface) *Controller {
	return &Controller{
		Service: service,
	}
}
