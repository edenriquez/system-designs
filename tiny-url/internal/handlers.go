package internal

import (
	controller "github.com/edenriquez/system-designs/tiny-url/pkg/url"
	"github.com/kataras/iris"
)

func healthCheckHandler(ctx iris.Context) {
	// check services healthy
	ctx.JSON(iris.Map{"result": "success"})
}

func urlEncodeHandler(ctx iris.Context) {
	url := ctx.PostValue("url")
	shortURL := controller.URLShortener(url)
	ctx.JSON(iris.Map{"url_shortened": shortURL})
}
