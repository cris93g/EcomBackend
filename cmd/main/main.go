package main

import (
	"github.com/cris93g/backendEcom/pkg/db"
	"github.com/cris93g/backendEcom/pkg/routes"
	"github.com/gofiber/fiber"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)



func main() {
	app := fiber.New()
	db.InitDatabase()
	defer db.DBConn.Close()
	routes.Routes(app)
	app.Listen(3001)
}