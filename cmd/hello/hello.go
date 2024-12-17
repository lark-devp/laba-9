package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Структура для приема JSON данных
type Response struct {
	Message string `json:"message"`
}

func helloHandler(c echo.Context) error {
	response := Response{Message: "Hello, web!"}
	return c.JSON(http.StatusOK, response)
}

func main() {
	// Создание нового экземпляра Echo
	e := echo.New()

	// Логирование
	e.Logger.SetHeader("${time_rfc3339} ${remote_ip} ${method} ${url} ${status}")

	// Роутинг
	e.GET("/get", helloHandler)

	// Обработка ошибок
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}

		c.JSON(code, map[string]interface{}{
			"error": err.Error(),
		})
	}

	// Запуск сервера
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
