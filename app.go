package fiber

// ... existing code ...

func (app *App) RequestHandler(ctx *fasthttp.RequestCtx) {
	c := app.AcquireCtx(ctx)
	defer app.ReleaseCtx(c)

	// Handle request
	_ = app.next(c)
}

// ... existing code ...