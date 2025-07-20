package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type FiberServer struct {
	*fiber.App
}

func NewFiber(env *Env) *FiberServer {
	app := fiber.New(fiber.Config{
		AppName:      env.SERVER_APP_NAME,
		ServerHeader: env.SERVER_SERVER_NAME,
		Prefork:      env.SERVER_PREFORK,
	})

	app.Use(logger.New())

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Server", env.SERVER_SERVER_NAME)
		return c.Next()
	})

	app.Use(func(c *fiber.Ctx) error {
		c.Set("X-Content-Type-Options", "nosniff")
		return c.Next()
	})

	app.Use(helmet.New(helmet.Config{
		XSSProtection:             "1; mode=block",
		ContentTypeNosniff:        "nosniff",
		XFrameOptions:             "SAMEORIGIN",
		ReferrerPolicy:            "no-referrer",
		CrossOriginEmbedderPolicy: "require-corp",
		CrossOriginOpenerPolicy:   "same-origin",
		CrossOriginResourcePolicy: "same-origin",
		OriginAgentCluster:        "?1",
		XDNSPrefetchControl:       "off",
		XDownloadOptions:          "noopen",
		XPermittedCrossDomain:     "none",
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins:     env.ALLOWED_ORIGINS,
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		ExposeHeaders:    "Content-Length, X-Requested-With",
		AllowCredentials: true,
		MaxAge:           300,
	}))

	app.Use(recover.New())

	app.Use(healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		LivenessEndpoint: "/live",
		ReadinessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		ReadinessEndpoint: "/ready",
	}))

	return &FiberServer{App: app}
}
