package middleware

import (
	"regexp"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

var (
	Compress = compress.New(
		compress.Config{
			Level: compress.LevelBestSpeed,
		})

	BasicAuth = basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin": "admin",
		},
		Next: filterURL,
	})

	Helmet = helmet.New()

	Limiter = limiter.New(limiter.Config{
		Max:               10,
		Expiration:        1 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	})

	Monitor = monitor.New(monitor.Config{Title: "Performance Metrics"})

	Logger = logger.New(logger.Config{
		// For more options, see the Config section
		TimeFormat: time.RFC3339Nano,
		TimeZone:   "Europe/Czechia",
		Format:     "[${time}] ${ip}:${port} ${pid} ${status} ${method} ${path}\n",
	})
)

func filterURL(c *fiber.Ctx) bool {
	originalURL := strings.ToLower(c.OriginalURL())

	protectedURLs := []*regexp.Regexp{
		regexp.MustCompile("^/metrics$"),
	}

	for _, url := range protectedURLs {
		if url.MatchString(originalURL) {
			return false
		}
	}
	return true
}
