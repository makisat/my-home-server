package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	_ "github.com/mattn/go-sqlite3"
)

type Link struct {
	Id		uint	`json:"id"`
	Name	string	`json:"name"`
	Url		string	`json:"url"`
}

func main() {
	// connect to database
	connectDB()
	defer db.Close()

	// create table if not exist
	initStatement := `CREATE TABLE IF NOT EXISTS links (id INTEGER PRIMARY KEY, link_name TEXT NOT NULL, url TEXT NOT NULL);`
	_, err := db.Exec(initStatement)
	if err != nil {
		log.Println("error executing initStatement: " + err.Error())
	}

	// new app
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, DELETE",
	}))

	// gets
	app.Get("/ping", ping)
	app.Get("/", get_links)

	// posts
	app.Post("/add-link", add_link)
	app.Delete("/delete-link/:id", delete_link)

	log.Fatal(app.Listen(":8080"))
}
