package rest

import (
	"time"

	"botperational/config"
	"botperational/internal/pkg/custommiddleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Fiber struct {
	*fiber.App
	Conf *config.Config `inject:"config"`
}

func (f *Fiber) Startup() error {
	// starting http
	f.App = fiber.New(fiber.Config{
		ReadTimeout:  time.Duration(f.Conf.Http.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(f.Conf.Http.WriteTimeout) * time.Second,
		ErrorHandler: custommiddleware.ErrorHandler(CodeMap, StatusMap),
	})

	// Middleware
	f.Use(recover.New(recover.Config{
		EnableStackTrace:  true,
		StackTraceHandler: custommiddleware.RecoverHandler,
	}))
	f.Use(requestid.New())
	f.Use(custommiddleware.Log(CodeMap, StatusMap))

	// CORS config
	f.Use(cors.New(cors.Config{
		AllowOrigins: "https://api-docs.runsystemdev.com",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	return nil
}

func (f *Fiber) Shutdown() error { return f.App.Shutdown() }
