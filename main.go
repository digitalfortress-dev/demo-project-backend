package main

import (
	"fmt"
	"main/database"
	"main/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	api := echo.New()
	api.Use(middleware.Logger())
	api.Use(middleware.Recover())

	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.POST, echo.GET},
	}))
	db := database.GetInstance()
	database.CheckTableInDatabase(db)
	m := database.GetMigrations(db)
	err := m.Migrate()
	if err == nil {
		fmt.Println("Migrations ran successfully.")
	} else {
		fmt.Println("Migrations failed.", err)
	}
	routes.DefineApiRoute(api)
	fmt.Println(api.Start(":1323"))
}
