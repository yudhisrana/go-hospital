package patient_handler

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/yudhisrana/go-hospital/internal/application/patient/dto"
	"github.com/yudhisrana/go-hospital/internal/domain/patient/valueobject"
	"github.com/yudhisrana/go-hospital/pkg/response"
)

func (h *Handler) CreatePatientHandler(c fiber.Ctx) error {
	var req dto.RequestPayload

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	result, err := h.usc.CreatePatient.Execute(c.Context(), &req)

	if err != nil {
		switch {
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

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return response.JsonSuccessResponse(c, fiber.StatusCreated, result, "Pasien berhasil dibuat")
}
