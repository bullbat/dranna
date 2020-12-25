package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.Text("hello world!")
	})
	app.Run(iris.Addr(":8080"))
}
