package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/Woddell/go-rest-api-course/internal/transport/http"
)

// App - Struct which contains things pointers to db connections, etc.
type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting Up Our App")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to setup server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	app := App{}

	if err := app.Run(); err != nil {
		fmt.Println("Error starting up REST API")
		fmt.Println(err)
	}
}
