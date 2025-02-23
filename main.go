package main

import (
	"fmt"
	"github.com/abdukarimxalilov/go-fiber/database"
	"github.com/abdukarimxalilov/go-fiber/handler"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App){
	app.Get("/api/v1/lead", handler.GetLeads)
	app.Get("/api/v1/lead/:id", handler.GetLead)
	app.Post("/api/v1/lead", handler.NewLead)
	app.Delete("/api/v1/lead/:id", handler.DeleteLead)

}

func initDatabase(){
	var err error 
	database.DBconn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil{
		panic("failed to connect database")
	}
	fmt.Println("Connection opened to database.")
	database.DBconn.AutoMigrate(&handler.Lead{})
	fmt.Println("database migrated.")
}


func main(){
	app := fiber.New()
	initDatabase()
	setupRoutes(app)

	app.Listen(3000)

	defer database.DBconn.DB().Close()
}