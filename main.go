package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var (
	listen   = os.Getenv("LISTEN")
	smtpUser = os.Getenv("SMTP_USER")
	smtpPass = os.Getenv("SMTP_PASS")
	smtpHost = os.Getenv("SMTP_HOST")
)

func email(subject, body string, to []string) error {
	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)
	return smtp.SendMail(smtpHost+":587", auth, smtpUser, to, []byte("Subject: "+subject+"\n"+"From: noreply@radixproject.org\n"+body))
}

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://radixproject.org",
	}))

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	app.Post("/form", func(c *fiber.Ctx) error {
		name := c.FormValue("name")

		// Get body as JSON
		var body map[string]string
		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error":   true,
				"message": "Invalid JSON body",
			})
		}

		out := fmt.Sprintf("Received submission for %s form\n\n", name)
		for k, v := range body {
			out += fmt.Sprintf("%s:\n%s\n\n", k, v)
		}
		if err := email(name+" form submission", out, []string{"info@radixproject.org"}); err != nil {
			log.Println(err)
			return c.Status(500).JSON(fiber.Map{
				"error":   true,
				"message": "Failed to submit form",
			})
		}

		return c.JSON(fiber.Map{
			"error":   false,
			"message": "OK",
		})
	})

	log.Fatal(app.Listen(listen))
}
