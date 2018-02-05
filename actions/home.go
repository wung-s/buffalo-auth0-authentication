package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.JSON(map[string]string{"message": "Welcome to Buffalo!"}))
}

// MySecuredEndpoint is mapped to `/info` route
func MySecuredEndpoint(c buffalo.Context) error {
	return c.Render(200, r.JSON("This is a secured endpoint"))
}
