package routes

import (
	"github.com/cris93g/backendEcom/pkg/controllers"
	"github.com/gofiber/fiber"
)


func Routes(r *fiber.App){
	r.Get("/api/v1/items",controllers.GetItems)
	r.Get("/api/v1/item/:id",controllers.GetItem)
	r.Post("/api/v1/items",controllers.AddToItems)
	r.Delete("/api/v1/item/:id",controllers.DeleteItem)
}