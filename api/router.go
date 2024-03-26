package api

import (
	"rent-car/api/handler"
	"rent-car/storage"
	"github.com/gin-gonic/gin"
)


// New ...
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.

func New(store storage.IStorage) *gin.Engine {
	h := handler.NewStrg(store)

	r := gin.Default()

	r.POST("/car", h.CreateCar)
	r.GET("/car/:id", h.GetAllCars)
	r.GET("/cars", h.GetAllCars)
	r.PUT("/car/:id", h.UpdateCar)
	r.DELETE("/car/:id", h.DeleteCar)
	
	r.POST("/customer", h.CreateCustomer)
	r.GET("/customer/:id", h.GetAllCustomer)
	r.GET("/customers", h.GetAllCustomer)
	r.PUT("/customer/:id", h.UpdateCustomer)
	r.DELETE("/customer/:id", h.DeleteCustomer)
	r.GET("/customercars/:id", h.GetAllCustomerCars)
	r.GET("/customercars",h.GetAllCustomerCars)


	r.POST("/order", h.CreateOrder)
	r.GET("/order/:id", h.GetAllOrder)
	r.GET("/orders", h.GetAllOrder)
	r.PUT("/order/:id", h.UpdateOrder)
	r.DELETE("/order/:id", h.DeleteOrder)

	return r
}
