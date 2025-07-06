package controller

import (
	"github.com/PreethiS-10/fiber-crud/internal/config"
	"github.com/PreethiS-10/fiber-crud/internal/models"
	"github.com/PreethiS-10/fiber-crud/internal/service"
	"github.com/gofiber/fiber/v2"
)

type ScheduleController struct {
    scheduleService service.ScheduleService
}

func NewScheduleController(scheduleService service.ScheduleService) *ScheduleController {
    return &ScheduleController{scheduleService: scheduleService}
}
func (s *ScheduleController) GetSchedules(c *fiber.Ctx) error {
	var schedules []models.Schedule
	if err := config.DB.Preload("Doctor").Find(&schedules).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve schedules"})
	}
	return c.JSON(schedules)
}

func (s *ScheduleController) GetSchedule(c *fiber.Ctx) error {
	id := c.Params("id")
	var schedule models.Schedule
	if err := config.DB.Preload("Doctor").First(&schedule, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Schedule not found"})
	}
	return c.JSON(schedule)
}

func (s *ScheduleController) CreateSchedule(c *fiber.Ctx) error {
	var schedule models.Schedule
	if err := c.BodyParser(&schedule); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := config.DB.Create(&schedule).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create schedule"})
	}
	return c.Status(201).JSON(schedule)
}

func (s *ScheduleController) UpdateSchedule(c *fiber.Ctx) error {
	id := c.Params("id")
	var schedule models.Schedule
	if err := config.DB.First(&schedule, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Schedule not found"})
	}
	if err := c.BodyParser(&schedule); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := config.DB.Save(&schedule).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update schedule"})
	}
	return c.JSON(schedule)
}

func (s *ScheduleController) DeleteSchedule(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.Schedule{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete schedule"})
	}
	return c.SendStatus(204)
}
