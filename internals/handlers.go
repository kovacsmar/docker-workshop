package internals

import (
	"context"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/kovacsmar/docker-workshop/views"
)

const appTimeout = time.Second * 10

func render(ctx *gin.Context, status int, template templ.Component) error {
	ctx.Status(status)
	return template.Render(ctx.Request.Context(), ctx.Writer)
}

func (app *Config) indexPageHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		defer cancel()

		render(ctx, http.StatusOK, views.Index())
	}
}

func (app *Config) getJokeHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		defer cancel()

		joke, err := app.getYoMammaJoke()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		render(ctx, http.StatusOK, views.BlockQuote((*views.YoMamaJoke)(joke)))
	}
}
