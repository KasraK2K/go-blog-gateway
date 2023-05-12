package routes

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		company := "embargo"

		requestURL := fmt.Sprintf("http://localhost:8000/v1/%s", company)
		req, err := http.NewRequest(http.MethodGet, requestURL, nil)
		if err != nil {
			fmt.Println("error while creating req", err)
			os.Exit(1)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("error making http get request error :", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		buf := new(strings.Builder)
		io.Copy(buf, resp.Body)

		c.Format(buf)
		return c.SendString(buf.String())
	})

	app.Get("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		company := "embargo"

		requestURL := fmt.Sprintf("http://localhost:8000/v1/%s/%s", company, id)
		req, err := http.NewRequest(http.MethodGet, requestURL, nil)
		if err != nil {
			fmt.Println("error while creating req", err)
			os.Exit(1)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("error making http get request error :", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		buf := new(strings.Builder)
		io.Copy(buf, resp.Body)

		c.Format(buf)
		return c.SendString(buf.String())
	})
}
