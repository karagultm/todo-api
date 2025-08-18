package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// var db = []Todo{{1, "Github'a projelerini yükle!", false}, {2, "Udemy Kursunu Bitir.", true}}

func main() {
	e := echo.New()
	db, err := gorm.Open(sqlite.Open("C:/Users/Administrator/Desktop/staj/go exercises/todoapi/database.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Todo{})

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	e.GET("/todos", listAll(db))
	e.POST("/todos", addTodo(db))
	e.PUT("/todos/:id", toggle(db))
	e.DELETE("/todos/:id", delete(db))

	e.Logger.Fatal(e.Start(":1323"))
}

func listAll(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var todos []Todo
		db.Find(&todos)
		return c.JSON(http.StatusOK, todos)
	}

}

func addTodo(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		todo := new(Todo)
		if err := c.Bind(todo); err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		// db = append(db, *todo)
		if err := db.Create(todo).Error; err != nil {
			return c.String(http.StatusInternalServerError, "cound not save the todo")
		}

		return c.JSON(http.StatusOK, todo)

	}
}
func toggle(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		param := c.Param("id")
		id, err := strconv.Atoi(param)
		if err != nil {
			return c.String(http.StatusBadRequest, "invalid id format")
		}
		var todo Todo
		if err := db.First(&todo, id).Error; err != nil {
			return c.String(http.StatusNotFound, "todo not found")
		}

		todo.Done = !todo.Done

		if err := db.Save(&todo).Error; err != nil {
			return c.String(http.StatusInternalServerError, " database coundnt save")
		}

		return c.JSON(http.StatusOK, todo)
	}
}
func delete(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "invalid id format")

		}

		result := db.Delete(&Todo{}, id)

		if err := result.RowsAffected; err == 0 {
			return c.String(http.StatusNotFound, "id not found on the list")
		}
		if err := result.Error; err != nil {
			return c.String(http.StatusInternalServerError, " database error")
		}

		return c.JSON(http.StatusOK, "todo deleted succesfuly")

	}
}
