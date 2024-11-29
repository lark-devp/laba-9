package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Структура для JSON-ответа
type Response struct {
	Message string `json:"message"` // Исправлено на правильный синтаксис для JSON-тегов
}

func handler(c echo.Context) error {
	name := c.QueryParam("name")

	// Проверка, было ли передано имя
	if name == "" {
		return c.JSON(http.StatusBadRequest, Response{Message: "Пожалуйста, введите ваше имя с помощью параметра 'name'."})
	}

	response := Response{Message: "Hello, " + name + "!"}
	return c.JSON(http.StatusOK, response)
}

func main() {
	// Создание нового экземпляра Echo
	e := echo.New()

	// Логирование
	e.Logger.SetHeader("${time_rfc3339} ${remote_ip} ${method} ${url} ${status}")

	// Роутинг
	e.GET("/api/user", handler)

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
	if err := e.Start(":9000"); err != nil {
		e.Logger.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
