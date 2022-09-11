package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/tus/tusd/pkg/filestore"

	tus "moobert/fiber-tus"

	tusd "github.com/tus/tusd/pkg/handler"
)

func main() {
	app := fiber.New()

	err := tus.Mount("/api/v1/file", app, newTusdConfig())
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(app.Listen(":3000"))
}

func newTusdConfig() tusd.Config {
	store := filestore.FileStore{
		Path: "./examples/uploads",
	}

	composer := tusd.NewStoreComposer()
	store.UseIn(composer)

	return tusd.Config{
		StoreComposer: composer,
	}
}
