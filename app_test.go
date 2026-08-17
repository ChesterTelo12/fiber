package fiber

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2/utils"
)

// ... existing code ...

func Test_Ctx_Locals_Leak_On_Error(t *testing.T) {
	app := New()

	app.Get("/error", func(c *Ctx) error {
		c.Locals("secret", "leak-data")
		return errors.New("something went wrong")
	})

	app.Get("/clean", func(c *Ctx) error {
		secret := c.Locals("secret")
		if secret != nil {
			t.Errorf("State leak detected! Found 'secret': %v", secret)
		}
		return c.SendString("ok")
	})

	// 1. Trigger the error route to populate and recycle the context
	resp, err := app.Test(httptest.NewRequest("GET", "/error", nil))
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, 500, resp.StatusCode)

	// 2. Trigger the clean route and verify no state leaked
	resp, err = app.Test(httptest.NewRequest("GET", "/clean", nil))
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, 200, resp.StatusCode)
}

// ... existing code ...