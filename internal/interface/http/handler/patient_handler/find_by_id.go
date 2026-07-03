package patient_handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/yudhisrana/go-hospital/pkg/response"
)

func (h *Handler) FindPatientByIDHandler(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return response.JsonErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	result, err := h.usc.FindPatientByID.Execute(c.Context(), id)
	if err != nil {
		return response.JsonErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	if result == nil {
		return response.JsonErrorResponse(c, fiber.StatusNotFound, "patient not found")
	}

	return response.JsonSuccessResponse(c, fiber.StatusOK, result, "Pasien berhasil ditemukan")
}



