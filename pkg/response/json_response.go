package response

import "github.com/gofiber/fiber/v3"

type SuccessResponse[T any] struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Data    T      `json:"data"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func JsonSuccessResponse[T any](c fiber.Ctx, code int, data T, message string) error {
	return c.Status(code).JSON(&SuccessResponse[T]{
		Code:    code,
		Status:  "success",
		Data:    data,
		Message: message,
	})
}

func JsonErrorResponse(c fiber.Ctx, code int, message string) error {
	return c.Status(code).JSON(&ErrorResponse{
		Code:    code,
		Status:  "error",
		Message: message,
	})
}
