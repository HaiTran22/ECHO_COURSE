package main

import (
	// "log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// đầu tiên là đky cái server
	server := echo.New()

	server.Use(middleware.Logger())
	// đky Route: GET, Path: "/" , errFunc: hello
	server.GET("/", hello)
	server.GET("/page2", page2)
	server.POST("/login", login, middleware.BasicAuth(basicAuth))

	// Chạy trên localhost:8080/
	server.Logger.Fatal(server.Start(":8080"))
}

// hàm hello trả về echo.Context
func hello(c echo.Context) error {
	// return trả về string "...", StatusOK là trả về 200
	return c.String(http.StatusOK, "Hello world!!!")
}

func page2(c echo.Context) error {
	return c.String(http.StatusOK, "Page 2")
}

func basicAuth(username string, password string, c echo.Context) (bool, error) {
	if username != "admin" || password != "123123" {
		return false, nil
	}

	return true, nil
}

type loginRequest struct {
	UserName string `json:"username" form:"username" xml:"username" query:"username" `
	Password string `json:"password" form:"password" xml:"password" query:"password" `
}
