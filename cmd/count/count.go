package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var counter int = 0

// Структура для JSON-ответов
type CountResponse struct {
	Count int `json:"count"`
}

type CountRequest struct {
	Count int `json:"count"`
}

func countHandler(c echo.Context) error {
	switch c.Request().Method {
	case http.MethodGet:
		return c.JSON(http.StatusOK, CountResponse{Count: counter})
	case http.MethodPost:
		var req CountRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "это не число"})
		}
		counter += req.Count
		return c.JSON(http.StatusOK, map[string]string{"message": "Значение счётчика обновлено"})
	default:
		return c.JSON(http.StatusMethodNotAllowed, map[string]string{"error": "Метод не поддерживается"})
	}
}

func main() {
	e := echo.New()

	// Добавление middleware для логирования
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Определение маршрута
	e.GET("/count", countHandler)
	e.POST("/count", countHandler)

	// Запуск сервера
	if err := e.Start(":3333"); err != nil {
		e.Logger.Fatal("Ошибка запуска сервера:", err)
	}
}
