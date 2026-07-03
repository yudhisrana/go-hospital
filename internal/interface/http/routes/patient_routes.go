package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/yudhisrana/go-hospital/internal/interface/http/handler/patient_handler"
)

func RegisterPatientRoutes(app *fiber.App, handler *patient_handler.Handler) {
	patientGroup := app.Group("/patients")

	patientGroup.Post("", handler.CreatePatientHandler)
	patientGroup.Put("/:id", handler.UpdatePatientHandler)
	patientGroup.Delete("/:id", handler.DeletePatientHandler)
	patientGroup.Get("/:id", handler.FindPatientByIDHandler)
	patientGroup.Get("", handler.FindAllPatientHandler)
}
