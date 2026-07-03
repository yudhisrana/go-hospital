package visit_handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/yudhisrana/go-hospital/internal/application/visit/dto"
	"github.com/yudhisrana/go-hospital/internal/application/visit/usecase"
	"github.com/yudhisrana/go-hospital/internal/domain/visit/entity"
	"github.com/yudhisrana/go-hospital/pkg/response"
)

type Handler struct {
	usc *usecase.VisitUseCase
}

func NewVisitHandler(usc *usecase.VisitUseCase) *Handler {
	return &Handler{usc: usc}
}

// POST /visits/register
func (h *Handler) RegisterVisitHandler(c fiber.Ctx) error {
	var req dto.RegisterVisitRequest

	if err := c.Bind().Body(&req); err != nil {
		return response.JsonErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	patientID, err := uuid.Parse(req.PatientID)
	if err != nil {
		return response.JsonErrorResponse(c, fiber.StatusBadRequest, "invalid patient_id")
	}

	visit, err := h.usc.RegisterVisit.Execute(c.Context(), patientID, req.Symptoms)
	if err != nil {
		return response.JsonErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.JsonSuccessResponse(c, fiber.StatusCreated, dto.NewVisitResponse(visit), "Pasien berhasil didaftar, nomor antrean: "+string(rune(visit.QueueNumber())))
}

// PATCH /visits/:id/examine
func (h *Handler) ExaminePatientHandler(c fiber.Ctx) error {
	visitID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.JsonErrorResponse(c, fiber.StatusBadRequest, "invalid visit_id")
	}

	var req dto.ExaminePatientRequest
	if err := c.Bind().Body(&req); err != nil {
		return response.JsonErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	prescriptions := make([]entity.Prescription, len(req.Prescriptions))
	for i, p := range req.Prescriptions {
		prescriptions[i] = entity.Prescription{
			Medicine: p.Medicine,
			Dosage:   p.Dosage,
			Quantity: p.Quantity,
		}
	}

	visit, err := h.usc.ExaminePatient.Execute(c.Context(), visitID, req.Diagnosis, prescriptions)
	if err != nil {
		if err == entity.ErrVisitNotFound {
			return response.JsonErrorResponse(c, fiber.StatusNotFound, "visit not found")
		}
		return response.JsonErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.JsonSuccessResponse(c, fiber.StatusOK, dto.NewVisitResponse(visit), "Pasien berhasil diperiksa, menunggu apotek")
}

// PATCH /visits/:id/dispense
func (h *Handler) DispenseMedicineHandler(c fiber.Ctx) error {
	visitID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.JsonErrorResponse(c, fiber.StatusBadRequest, "invalid visit_id")
	}

	visit, err := h.usc.DispenseMedicine.Execute(c.Context(), visitID)
	if err != nil {
		if err == entity.ErrVisitNotFound {
			return response.JsonErrorResponse(c, fiber.StatusNotFound, "visit not found")
		}
		return response.JsonErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.JsonSuccessResponse(c, fiber.StatusOK, dto.NewVisitResponse(visit), "Obat berhasil diberikan, visit selesai")
}

// GET /visits
func (h *Handler) FindAllVisitHandler(c fiber.Ctx) error {
	visits, err := h.usc.FindAllVisit.Execute(c.Context())
	if err != nil {
		return response.JsonErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	if visits == nil {
		visits = make([]*entity.Visit, 0)
	}

	responses := make([]*dto.VisitResponse, len(visits))
	for i, v := range visits {
		responses[i] = dto.NewVisitResponse(v)
	}

	return response.JsonSuccessResponse(c, fiber.StatusOK, responses, "Data visit berhasil diambil")
}

// GET /visits/:id
func (h *Handler) FindVisitByIDHandler(c fiber.Ctx) error {
	visitID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.JsonErrorResponse(c, fiber.StatusBadRequest, "invalid visit_id")
	}

	visit, err := h.usc.FindVisitByID.Execute(c.Context(), visitID)
	if err != nil {
		return response.JsonErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	if visit == nil {
		return response.JsonErrorResponse(c, fiber.StatusNotFound, "visit not found")
	}

	return response.JsonSuccessResponse(c, fiber.StatusOK, dto.NewVisitResponse(visit), "Visit berhasil ditemukan")
}

// GET /visits?status=waiting_doctor
func (h *Handler) FindVisitByStatusHandler(c fiber.Ctx) error {
	status := c.Query("status", "")
	if status == "" {
		return response.JsonErrorResponse(c, fiber.StatusBadRequest, "status parameter required")
	}

	visits, err := h.usc.FindVisitByStatus.Execute(c.Context(), status)
	if err != nil {
		return response.JsonErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	if visits == nil {
		visits = make([]*entity.Visit, 0)
	}

	responses := make([]*dto.VisitResponse, len(visits))
	for i, v := range visits {
		responses[i] = dto.NewVisitResponse(v)
	}

	return response.JsonSuccessResponse(c, fiber.StatusOK, responses, "Visit berhasil diambil berdasarkan status")
}
