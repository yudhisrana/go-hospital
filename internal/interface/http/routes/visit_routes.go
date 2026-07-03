package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/yudhisrana/go-hospital/internal/interface/http/handler/visit_handler"
)

func RegisterVisitRoutes(app *fiber.App, handler *visit_handler.Handler) {
	visitGroup := app.Group("/visits")

	// Registration flow
	visitGroup.Post("/register", handler.RegisterVisitHandler)

	// Doctor examination
	visitGroup.Patch("/:id/examine", handler.ExaminePatientHandler)

	// Pharmacy dispensing
	visitGroup.Patch("/:id/dispense", handler.DispenseMedicineHandler)

	// View visits
	visitGroup.Get("", handler.FindAllVisitHandler)
	visitGroup.Get("/:id", handler.FindVisitByIDHandler)

	// Filter by status
	visitGroup.Get("/status/:status", handler.FindVisitByStatusHandler)
}
