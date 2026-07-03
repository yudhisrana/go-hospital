package patient_handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/yudhisrana/go-hospital/pkg/response"
)

func (h *Handler) DeletePatientHandler(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return response.JsonErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	if err := h.usc.DeletePatient.Execute(c.Context(), id); err != nil {
		return response.JsonErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.JsonSuccessResponse(c, fiber.StatusNoContent, struct{}{}, "Pasien berhasil dihapus")
}





