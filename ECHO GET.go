package main

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo"
)

func getStud(c echo.Context) error {
	name := c.QueryParam("name")
	kurs := c.QueryParam("kurs")
	return c.String(http.StatusOK, fmt.Sprintf("Student name: %s\nand kurs:%s\n", name, kurs))
}

func qq(c echo.Context) error {

	return c.String(http.StatusOK, "qq")
}
func main() {
	fmt.Println("Hello")

	e := echo.New()

	e.GET("/", qq)
	e.GET("/student", getStud)
	e.Logger.Fatal(e.Start(":1323"))
}

// C:\Go\bin\go.exe run C:\Users\Работа\Desktop\go\GOLANG.go
