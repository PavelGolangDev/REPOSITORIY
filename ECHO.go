package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type Student struct {
	name string `json:"name"`
	kurs string `json:"kurs"`
}

func getStud(c echo.Context) error {
	name := c.QueryParam("name")
	kurs := c.QueryParam("kurs")
	return c.String(http.StatusOK, fmt.Sprintf("Student name: %s\nand kurs:%s\n", name, kurs))
}

func addStudent(c echo.Context) error {
	student := Student{}

	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the body:%s", err)
		return c.String(http.StatusInternalServerError, " ")
	}
	err = json.Unmarshal(b, &student)
	if err != nil {
		log.Printf("Failed unmarshaling:%s", err)
		return c.String(http.StatusInternalServerError, " ")
	}
	log.Printf("This is student:%#v", student)
	return c.String(http.StatusOK, "we got student")
}
func main() {
	e := echo.New()

	e.GET("/student", getStud)
	e.Logger.Fatal(e.Start(":1323"))
	e.POST("/addstudent", addStudent)
}

// C:\Go\bin\go.exe run C:\Users\Работа\Desktop\go\GOLANG.go
