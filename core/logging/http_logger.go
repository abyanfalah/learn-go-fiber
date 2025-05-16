package logging

import (
	"fmt"
	"log"

	"slices"

	"github.com/gofiber/fiber/v2"
)

// func HttpLogger(c *fiber.Ctx) error {
// 	var reqBody string
// 	if c.Request().Body() != nil {
// 		reqBody = string(c.Body())
// 		fmt.Printf("req body: %s", reqBody)
// 	}

// 	bodyCopy := slices.Clone(c.Response().Body())
// 	if len(bodyCopy) > 0 {
// 		fmt.Printf("res body: %s", string(bodyCopy))
// 	}

// 	if err := c.Next(); err != nil {
// 		return err
// 	}

// 	return nil
// }

func HttpLogger(c *fiber.Ctx) error {
	// fmt.Printf("%s %s\n", c.Method(), c.OriginalURL())
	//
	log.Println("=== HTTP REQUEST ===")
	fmt.Printf("%-14s: %s\n", "Method", c.Method())
	fmt.Printf("%-14s: %s\n", "Path", c.Path())
	fmt.Printf("%-14s: %s\n", "Address", c.IP())
	if c.Body() != nil {
		fmt.Printf("%-14s: %s", "Request body", string(c.Body()))
	}

	if err := c.Next(); err != nil {
		return err
	}

	// Log response
	resBody := slices.Clone(c.Response().Body())
	fmt.Printf("%-14s: %d\n", "Response code", c.Response().StatusCode())
	if len(resBody) > 0 {
		fmt.Printf("%-14s: %s\n", "Response Body", string(resBody))
	}

	return nil
}

// // Intercept response
// var resBody bytes.Buffer

// c.Response().SetBodyStreamWriter(func(w *io.PipeWriter) {
// 	defer w.Close() // always close it to avoid hanging

// 	// 👇 This is where you write your response
// 	// You can modify, wrap, or intercept here

// 	body := c.Response().Body()
// 	_, _ = io.MultiWriter(w, logBuffer).Write(body)
// })
// // Call next middleware / handler
// err := c.Next()

// // After next handler runs, we can access response
// fmt.Printf("\n=== REQUEST ===\n")
// fmt.Printf("Method: %s\nURL: %s\nHeaders: %v\nBody: %s\n",
// 	c.Method(), c.OriginalURL(), c.GetReqHeaders(), reqBody)

// fmt.Printf("\n=== RESPONSE ===\n")
// fmt.Printf("Status: %d\nHeaders: %v\nBody: %s\n",
// 	c.Response().StatusCode(), c.Response().Header.Header(), resBody.String())

// return err
