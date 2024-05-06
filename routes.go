package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// test the connection
func ping(c *fiber.Ctx) error {
	return c.SendString("pong")
}

// get links from the database
func get_links(c *fiber.Ctx) error {
	rows, err := db.Query("SELECT * FROM links")
	if err != nil {
		log.Fatal("error at query" + err.Error())
		return c.SendStatus(http.StatusBadRequest)
	}

	defer rows.Close()

	var links []Link

	for rows.Next() {
		link := Link {}

		rows.Scan(&link.Id, &link.Name, &link.Url)

		log.Printf("id: %d", link.Id)
		log.Println("name: " + link.Name)
		log.Println("url: " + link.Url)

		links = append(links, link)
	}

	return c.JSON(links)
}

// add new link to the database
func add_link(c *fiber.Ctx) error {
	stmt, err := db.Prepare("INSERT INTO links (link_name, url) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	new_link := new(Link)
	if err := c.BodyParser(new_link); err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	log.Println("name: " + new_link.Name)
	log.Println("url: " + new_link.Url)

	_, err = stmt.Exec(new_link.Name, new_link.Url)
	if err != nil {
		log.Fatal(err)
	}

	return c.SendStatus(200)
}

func delete_link(c *fiber.Ctx) error {
	stmt, err := db.Prepare("DELETE FROM links WHERE id = ?")
	if err != nil {
		log.Fatal("error at prepare: " + err.Error())
		return c.SendStatus(http.StatusBadRequest)
	}
	defer stmt.Close()

	target_link := c.Params("id")
	log.Println("target link: " + target_link)

	_, err = stmt.Exec(target_link)
	if err != nil {
		log.Fatal("error at execute: " + err.Error())
		return c.SendStatus(http.StatusBadRequest)
	}

	return c.SendStatus(200)
}
