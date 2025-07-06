package controller

import (
    "net/http"
    "strconv"

    "github.com/PreethiS-10/fiber-crud/internal/models"
    "github.com/PreethiS-10/fiber-crud/internal/service"
    "github.com/gofiber/fiber/v2"
)

type PatientController struct {
    patientService service.PatientService
}

func NewPatientController(patientService service.PatientService) *PatientController {
    return &PatientController{patientService}
}

func (pc *PatientController) RegisterRoutes(app *fiber.App) {
    patient := app.Group("/patients")

    patient.Get("/", pc.GetPatients)
    patient.Get("/:id", pc.GetPatient)
    patient.Post("/", pc.CreatePatient)
    patient.Put("/:id", pc.UpdatePatient)
    patient.Delete("/:id", pc.DeletePatient)
}

// GET /patients
func (pc *PatientController) GetPatients(c *fiber.Ctx) error {
    patients, err := pc.patientService.GetAll()
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(patients)
}

// GET /patients/:id
func (pc *PatientController) GetPatient(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid patient ID"})
    }

    patient, err := pc.patientService.GetByID(uint(id))
    if err != nil {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "patient not found"})
    }
    return c.JSON(patient)
}

// POST /patients
func (pc *PatientController) CreatePatient(c *fiber.Ctx) error {
    var patient models.Patient
    if err := c.BodyParser(&patient); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
    }

    if err := pc.patientService.Create(&patient); err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.Status(http.StatusCreated).JSON(patient)
}

// PUT /patients/:id
func (pc *PatientController) UpdatePatient(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid patient ID"})
    }

    var patient models.Patient
    if err := c.BodyParser(&patient); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
    }

    patient.ID = uint(id) // Ensure ID matches path param

    if err := pc.patientService.Update(&patient); err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(patient)
}

// DELETE /patients/:id
func (pc *PatientController) DeletePatient(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid patient ID"})
    }

    if err := pc.patientService.Delete(uint(id)); err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.SendStatus(http.StatusNoContent)
}
