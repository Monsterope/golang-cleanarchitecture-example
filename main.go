package main

import "cleanarchitecture-example/app"

func main() {

	app := app.NewApp()

	app.Start(":8080")
}
