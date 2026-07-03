package patient_handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/yudhisrana/go-hospital/internal/application/patient/dto"
	"github.com/yudhisrana/go-hospital/pkg/response"
)

func (h *Handler) FindAllPatientHandler(c fiber.Ctx) error {
	result, err := h.usc.FindAllPatient.Execute(c.Context())
	if err != nil {
		return response.JsonErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	if result == nil {
		result = make([]*dto.ResponsePayload, 0)
	}

	return response.JsonSuccessResponse(c, fiber.StatusOK, result, "Data pasien berhasil diambil")
}





