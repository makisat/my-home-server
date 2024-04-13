package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

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
	initStatement := `CREATE TABLE IF NOT EXISTS links (link_id INTEGER PRIMARY KEY, link_name TEXT NOT NULL, url TEXT NOT NULL);`
	_, err := db.Exec(initStatement)
	if err != nil {
		log.Println("error executing initStatement: " + err.Error())
	}

	// new app
	app := fiber.New()

	// gets
	app.Get("/ping", ping)
	app.Get("/", get_links)

	// posts
	app.Post("/add-link", add_link)
	app.Post("/delete-link", delete_link)

	log.Fatal(app.Listen(":8080"))
}
