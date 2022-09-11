package tus

import (
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"

	tusd "github.com/tus/tusd/pkg/handler"
)

func Mount(prefix string, app *fiber.App, config tusd.Config) error {
	config.BasePath = prefix

	handler, err := tusd.NewHandler(config)
	if err != nil {
		return err
	}

	group := app.Group(
		prefix,
		adaptor.HTTPMiddleware(handler.Middleware),
	)

	group.Post("", adaptor.HTTPHandlerFunc(handler.PostFile))
	group.Head(":id", adaptor.HTTPHandlerFunc(handler.HeadFile))
	group.Patch(":id", adaptor.HTTPHandlerFunc(handler.PatchFile))

	if !config.DisableDownload {
		group.Get(":id", adaptor.HTTPHandlerFunc(handler.GetFile))
	}

	if config.StoreComposer.UsesTerminater && !config.DisableTermination {
		group.Delete(":id", adaptor.HTTPHandlerFunc(handler.DelFile))
	}

	return nil
}
