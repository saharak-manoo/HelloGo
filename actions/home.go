package actions

import "github.com/gobuffalo/buffalo"
import "fmt"

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	fmt.Println("Hello, Saharak")
	fmt.Println("Hello, 世界")

	return c.Render(200, r.JSON(map[string]string{"message": "Welcome to Buffalo! dove"}))
}
