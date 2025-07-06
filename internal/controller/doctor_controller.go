package controller

import (
	"github.com/PreethiS-10/fiber-crud/internal/config"
	"github.com/PreethiS-10/fiber-crud/internal/service"
	"github.com/PreethiS-10/fiber-crud/internal/models"
	"github.com/gofiber/fiber/v2"
)

type DoctorController struct {
	doctorService service.DoctorService
}

func NewDoctorController(doctorService service.DoctorService) *DoctorController {
	return &DoctorController{doctorService: doctorService}
}
// Methods on *DoctorController receiver
func (dc *DoctorController) GetDoctors(c *fiber.Ctx) error {
	var doctors []models.Doctor
	if err := config.DB.Preload("Department").Find(&doctors).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve doctors",
		})
	}
	return c.JSON(doctors)
}

func (dc *DoctorController) GetDoctor(c *fiber.Ctx) error {
	id := c.Params("id")
	var doctor models.Doctor
	if err := config.DB.Preload("Department").First(&doctor, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Doctor not found",
		})
	}
	return c.JSON(doctor)
}

func (dc *DoctorController) CreateDoctor(c *fiber.Ctx) error {
	var doctor models.Doctor
	if err := c.BodyParser(&doctor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}
	if err := config.DB.Create(&doctor).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create doctor",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(doctor)
}

func (dc *DoctorController) UpdateDoctor(c *fiber.Ctx) error {
	id := c.Params("id")
	var doctor models.Doctor
	if err := config.DB.First(&doctor, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Doctor not found",
		})
	}

	if err := c.BodyParser(&doctor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	if err := config.DB.Save(&doctor).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update doctor",
		})
	}
	return c.JSON(doctor)
}

func (dc *DoctorController) DeleteDoctor(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.Doctor{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete doctor",
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
