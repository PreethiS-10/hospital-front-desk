package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/PreethiS-10/fiber-crud/internal/models"
	"github.com/PreethiS-10/fiber-crud/internal/config"
	"github.com/PreethiS-10/fiber-crud/internal/service"
)

type AppointmentController struct{
	AppointmentService service.AppointmentService
}


func NewAppointmentController(AppointmentService service.AppointmentService) *AppointmentController {
    return &AppointmentController{AppointmentService}
}

func (a *AppointmentController) GetAppointments(c *fiber.Ctx) error {
	var appointments []models.Appointment
	if err := config.DB.Preload("Patient").Preload("Doctor").Find(&appointments).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve appointments"})
	}
	return c.JSON(appointments)
}

func (a *AppointmentController) GetAppointment(c *fiber.Ctx) error {
	id := c.Params("id")
	var appointment models.Appointment
	if err := config.DB.Preload("Patient").Preload("Doctor").First(&appointment, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Appointment not found"})
	}
	return c.JSON(appointment)
}

func (a *AppointmentController) CreateAppointment(c *fiber.Ctx) error {
	var appointment models.Appointment
	if err := c.BodyParser(&appointment); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := config.DB.Create(&appointment).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create appointment"})
	}
	return c.Status(201).JSON(appointment)
}

func (a *AppointmentController) UpdateAppointment(c *fiber.Ctx) error {
	id := c.Params("id")
	var appointment models.Appointment
	if err := config.DB.First(&appointment, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Appointment not found"})
	}
	if err := c.BodyParser(&appointment); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := config.DB.Save(&appointment).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update appointment"})
	}
	return c.JSON(appointment)
}

func (a *AppointmentController) DeleteAppointment(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.Appointment{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete appointment"})
	}
	return c.SendStatus(204)
}
