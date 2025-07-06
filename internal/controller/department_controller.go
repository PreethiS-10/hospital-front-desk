package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/PreethiS-10/fiber-crud/internal/models"
	"github.com/PreethiS-10/fiber-crud/internal/service"
)

type DepartmentController struct {
	DepartmentService service.DepartmentService
}

func NewDepartmentController(service service.DepartmentService) *DepartmentController {
	return &DepartmentController{DepartmentService: service}
}

func (dc *DepartmentController) GetDepartments(c *fiber.Ctx) error {
	departments, err := dc.DepartmentService.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve departments"})
	}
	return c.JSON(departments)
}

func (dc *DepartmentController) GetDepartment(c *fiber.Ctx) error {
	idParam := c.Params("id")
	var id uint
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid department ID"})
	}
	department, err := dc.DepartmentService.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Department not found"})
	}
	return c.JSON(department)
}

func (dc *DepartmentController) CreateDepartment(c *fiber.Ctx) error {
	var department models.Department
	if err := c.BodyParser(&department); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	created := dc.DepartmentService.Create(&department)
	return c.Status(fiber.StatusCreated).JSON(created)
}

func (dc *DepartmentController) UpdateDepartment(c *fiber.Ctx) error {
	idParam := c.Params("id")
	var id uint
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid department ID"})
	}
	var department models.Department
	if err := c.BodyParser(&department); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	department.ID = id
	updated := dc.DepartmentService.Update(&department)
	return c.JSON(updated)
}
func (dc *DepartmentController) DeleteDepartment(c *fiber.Ctx) error {
	idParam := c.Params("id")
	// Convert idParam from string to uint
	var id uint
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid department ID"})
	}
	if err := dc.DepartmentService.Delete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete department"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

