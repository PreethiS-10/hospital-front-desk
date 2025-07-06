package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/PreethiS-10/fiber-crud/internal/config"
	"github.com/PreethiS-10/fiber-crud/internal/models"
)

// -------------------- PATIENT --------------------
func GetPatients(c *fiber.Ctx) error {
	var patients []models.Patient
	if err := config.DB.Find(&patients).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve patients"})
	}
	return c.JSON(patients)
}

func GetPatient(c *fiber.Ctx) error {
	id := c.Params("id")
	var patient models.Patient
	if err := config.DB.First(&patient, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Patient not found"})
	}
	return c.JSON(patient)
}

func CreatePatient(c *fiber.Ctx) error {
	fmt.Println("CreatePatient endpoint hit")
	var patient models.Patient
	if err := c.BodyParser(&patient); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := config.DB.Create(&patient).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create patient"})
	}
	return c.Status(201).JSON(patient)
}

func UpdatePatient(c *fiber.Ctx) error {
	id := c.Params("id")
	var patient models.Patient
	if err := config.DB.First(&patient, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Patient not found"})
	}
	if err := c.BodyParser(&patient); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := config.DB.Save(&patient).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update patient"})
	}
	return c.JSON(patient)
}

func DeletePatient(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.Patient{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete patient"})
	}
	return c.SendStatus(204)
}

// -------------------- DOCTOR --------------------
func GetDoctors(c *fiber.Ctx) error {
	var doctors []models.Doctor
	if err := config.DB.Preload("Department").Find(&doctors).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve doctors"})
	}
	return c.JSON(doctors)
}

func GetDoctor(c *fiber.Ctx) error {
	id := c.Params("id")
	var doctor models.Doctor
	if err := config.DB.Preload("Department").First(&doctor, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Doctor not found"})
	}
	return c.JSON(doctor)
}

func CreateDoctor(c *fiber.Ctx) error {
	var doctor models.Doctor
	if err := c.BodyParser(&doctor); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := config.DB.Create(&doctor).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create doctor"})
	}
	return c.Status(201).JSON(doctor)
}

func UpdateDoctor(c *fiber.Ctx) error {
	id := c.Params("id")
	var doctor models.Doctor
	if err := config.DB.First(&doctor, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Doctor not found"})
	}
	if err := c.BodyParser(&doctor); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := config.DB.Save(&doctor).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update doctor"})
	}
	return c.JSON(doctor)
}

func DeleteDoctor(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.Doctor{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete doctor"})
	}
	return c.SendStatus(204)
}

// -------------------- DEPARTMENT --------------------
func GetDepartments(c *fiber.Ctx) error {
	var departments []models.Department
	if err := config.DB.Find(&departments).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve departments"})
	}
	return c.JSON(departments)
}

func GetDepartment(c *fiber.Ctx) error {
	id := c.Params("id")
	var department models.Department
	if err := config.DB.First(&department, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Department not found"})
	}
	return c.JSON(department)
}

func CreateDepartment(c *fiber.Ctx) error {
	var department models.Department
	if err := c.BodyParser(&department); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := config.DB.Create(&department).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create department"})
	}
	return c.Status(201).JSON(department)
}

func UpdateDepartment(c *fiber.Ctx) error {
	id := c.Params("id")
	var department models.Department
	if err := config.DB.First(&department, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Department not found"})
	}
	if err := c.BodyParser(&department); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := config.DB.Save(&department).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update department"})
	}
	return c.JSON(department)
}

func DeleteDepartment(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.Department{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete department"})
	}
	return c.SendStatus(204)
}

// -------------------- SCHEDULE --------------------
func GetSchedules(c *fiber.Ctx) error {
	var schedules []models.Schedule
	if err := config.DB.Preload("Doctor").Find(&schedules).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve schedules"})
	}
	return c.JSON(schedules)
}

func GetSchedule(c *fiber.Ctx) error {
	id := c.Params("id")
	var schedule models.Schedule
	if err := config.DB.Preload("Doctor").First(&schedule, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Schedule not found"})
	}
	return c.JSON(schedule)
}

func CreateSchedule(c *fiber.Ctx) error {
	var schedule models.Schedule
	if err := c.BodyParser(&schedule); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := config.DB.Create(&schedule).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create schedule"})
	}
	return c.Status(201).JSON(schedule)
}

func UpdateSchedule(c *fiber.Ctx) error {
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

func DeleteSchedule(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.Schedule{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete schedule"})
	}
	return c.SendStatus(204)
}

// -------------------- APPOINTMENT --------------------
func GetAppointments(c *fiber.Ctx) error {
	var appointments []models.Appointment
	if err := config.DB.Preload("Patient").Preload("Doctor").Find(&appointments).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve appointments"})
	}
	return c.JSON(appointments)
}

func GetAppointment(c *fiber.Ctx) error {
	id := c.Params("id")
	var appointment models.Appointment
	if err := config.DB.Preload("Patient").Preload("Doctor").First(&appointment, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Appointment not found"})
	}
	return c.JSON(appointment)
}

func CreateAppointment(c *fiber.Ctx) error {
	var appointment models.Appointment
	if err := c.BodyParser(&appointment); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := config.DB.Create(&appointment).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create appointment"})
	}
	return c.Status(201).JSON(appointment)
}

func UpdateAppointment(c *fiber.Ctx) error {
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

func DeleteAppointment(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.Appointment{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete appointment"})
	}
	return c.SendStatus(204)
}
