// package logging

// import (
// 	"fmt"
// 	"learn-fiber/core/exception"

// 	"github.com/gofiber/fiber/v2"
// )

// func HttpLogger(c *fiber.Ctx) error {

// 	// Clone request body
// 	var reqBody string
// 	if c.Request().Body() != nil {
// 		reqBody = string(c.Body())
// 		fmt.Printf("req body: %s", reqBody)
// 	}

// 	if c.Response().Body() != nil {
// 		fmt.Printf("res body: %s", string(c.Response().Body()))
// 	}

// 	if err := c.Next(); err != nil {
// 		return exception.Handle(err)
// 	}

// 	return nil

// 	// // Intercept response
// 	// var resBody bytes.Buffer

// 	// c.Response().SetBodyStreamWriter(func(w *io.PipeWriter) {
// 	// 	defer w.Close() // always close it to avoid hanging

// 	// 	// 👇 This is where you write your response
// 	// 	// You can modify, wrap, or intercept here

// 	// 	body := c.Response().Body()
// 	// 	_, _ = io.MultiWriter(w, logBuffer).Write(body)
// 	// })
// 	// // Call next middleware / handler
// 	// err := c.Next()

// 	// // After next handler runs, we can access response
// 	// fmt.Printf("\n=== REQUEST ===\n")
// 	// fmt.Printf("Method: %s\nURL: %s\nHeaders: %v\nBody: %s\n",
// 	// 	c.Method(), c.OriginalURL(), c.GetReqHeaders(), reqBody)

// 	// fmt.Printf("\n=== RESPONSE ===\n")
// 	// fmt.Printf("Status: %d\nHeaders: %v\nBody: %s\n",
// 	// 	c.Response().StatusCode(), c.Response().Header.Header(), resBody.String())

// 	// return err

// }
