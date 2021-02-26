package controllers

import (
	"github.com/cris93g/backendEcom/pkg/db"
	"github.com/cris93g/backendEcom/pkg/models"
	"github.com/gofiber/fiber"
)


func GetItems(c *fiber.Ctx) {
	db := db.DBConn
	var items []models.GoProducts
	db.Find(&items)
	c.JSON(items)
}


func AddToItems(c *fiber.Ctx){
	db:= db.DBConn
	item := new(models.GoProducts)
	if err := c.BodyParser(item); err != nil{
		c.Status(500).Send(err)
		return
	}
	db.Create(&item)
	c.JSON(item)
}

func GetItem(c *fiber.Ctx){
	id:= c.Params("id")
	db:= db.DBConn
	var item models.GoProducts
	db.Find(&item,id)
	if item.Name==""{
		c.Status(500).Send("could not find that specific item")
	}
	c.JSON(item)
}

func DeleteItem(c *fiber.Ctx){
	id:= c.Params("id")
	db:= db.DBConn
	var item models.GoProducts
	db.First(&item,id)
	if item.Name == ""{
		c.Status(500).Send("no item found with that id")
	}
	db.Delete(&item)
	c.Send("item is succesfully deleted")
}