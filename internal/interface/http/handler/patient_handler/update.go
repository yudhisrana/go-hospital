package patient_handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/yudhisrana/go-hospital/internal/application/patient/dto"
	"github.com/yudhisrana/go-hospital/internal/domain/patient/entity"
	"github.com/yudhisrana/go-hospital/internal/domain/patient/valueobject"
	"github.com/yudhisrana/go-hospital/pkg/response"
	"errors"
)

func (h *Handler) UpdatePatientHandler(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return response.JsonErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	var req dto.RequestPayload
	if err := c.Bind().Body(&req); err != nil {
		return response.JsonErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	result, err := h.usc.UpdatePatient.Execute(c.Context(), id, &req)

	if err != nil {
		switch {
		case errors.Is(err, entity.ErrPatientNotChange):
			return response.JsonErrorResponse(c, fiber.StatusBadRequest, entity.ErrPatientNotChange.Error())
		case errors.Is(err, valueobject.ErrorPatientNIKEmpty):
			return response.JsonErrorResponse(c, fiber.StatusBadRequest, valueobject.ErrorPatientNIKEmpty.Error())
		case errors.Is(err, valueobject.ErrorPatientNIKInvalid):
			return response.JsonErrorResponse(c, fiber.StatusBadRequest, valueobject.ErrorPatientNIKInvalid.Error())
		case errors.Is(err, valueobject.ErrorPatientNameEmpty):
			return response.JsonErrorResponse(c, fiber.StatusBadRequest, valueobject.ErrorPatientNameEmpty.Error())
		case errors.Is(err, valueobject.ErrorPatientNameInvalid):
			return response.JsonErrorResponse(c, fiber.StatusBadRequest, valueobject.ErrorPatientNameInvalid.Error())
		case errors.Is(err, valueobject.ErrorPatientAgeInvalid):
			return response.JsonErrorResponse(c, fiber.StatusBadRequest, valueobject.ErrorPatientAgeInvalid.Error())
		case errors.Is(err, valueobject.ErrorPatientGenderEmpty):
			return response.JsonErrorResponse(c, fiber.StatusBadRequest, valueobject.ErrorPatientGenderEmpty.Error())
		case errors.Is(err, valueobject.ErrorPatientGenderInvalid):
			return response.JsonErrorResponse(c, fiber.StatusBadRequest, valueobject.ErrorPatientGenderInvalid.Error())
		case errors.Is(err, valueobject.ErrorPatientBirthDateZero):
			return response.JsonErrorResponse(c, fiber.StatusBadRequest, valueobject.ErrorPatientBirthDateZero.Error())
		case errors.Is(err, valueobject.ErrorPatientBirthDateFuture):
			return response.JsonErrorResponse(c, fiber.StatusBadRequest, valueobject.ErrorPatientBirthDateFuture.Error())
		case errors.Is(err, valueobject.ErrorPatientAddressEmpty):
			return response.JsonErrorResponse(c, fiber.StatusBadRequest, valueobject.ErrorPatientAddressEmpty.Error())
		case errors.Is(err, valueobject.ErrorPatientPhoneEmpty):
			return response.JsonErrorResponse(c, fiber.StatusBadRequest, valueobject.ErrorPatientPhoneEmpty.Error())
		case errors.Is(err, valueobject.ErrorPatientPhoneInvalid):
			return response.JsonErrorResponse(c, fiber.StatusBadRequest, valueobject.ErrorPatientPhoneInvalid.Error())
		}
		return response.JsonErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	if result == nil {
		return response.JsonErrorResponse(c, fiber.StatusNotFound, "patient not found")
	}

	return response.JsonSuccessResponse(c, fiber.StatusOK, result, "Pasien berhasil diperbarui")
}

