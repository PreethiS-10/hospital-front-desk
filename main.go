package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/PreethiS-10/fiber-crud/internal/config"
	"github.com/PreethiS-10/fiber-crud/internal/models"
	"github.com/PreethiS-10/fiber-crud/internal/repository"
	"github.com/PreethiS-10/fiber-crud/internal/service"
	"github.com/PreethiS-10/fiber-crud/internal/controller"
	"github.com/PreethiS-10/fiber-crud/internal/router"
)

func main() {
	app := fiber.New()

	// Connect to DB
	config.ConnectDatabase()

	// Migrate models
	err := config.DB.AutoMigrate(
		&models.Patient{},
		&models.Doctor{},
		&models.Department{},
		&models.Schedule{},
		&models.Appointment{},
	)
	if err != nil {
		fmt.Println("Migration error:", err)
		return
	}

	// Initialize repositories
	patientRepo := repository.NewPatientRepository(config.DB)
	doctorRepo := repository.NewDoctorRepository(config.DB)
	departmentRepo := repository.NewDepartmentRepository(config.DB)
	scheduleRepo := repository.NewScheduleRepository(config.DB)
	appointmentRepo := repository.NewAppointmentRepository(config.DB)

	// Initialize services
	patientService := service.NewPatientService(patientRepo)
	doctorService := service.NewDoctorService(doctorRepo)
	departmentService := service.NewDepartmentService(departmentRepo)
	scheduleService := service.NewScheduleService(scheduleRepo)
	appointmentService := service.NewAppointmentService(appointmentRepo)

	// Initialize controllers
	patientController := controller.NewPatientController(patientService)
	doctorController := controller.NewDoctorController(doctorService) // Make sure NewDoctorController exists in the controller package; if not, implement it as shown below:

	// In internal/controller/doctor_controller.go, add:
	// package controller
	// import "github.com/PreethiS-10/fiber-crud/internal/service"
	// type DoctorController struct {
	//     Service service.DoctorService
	// }
	// func NewDoctorController(s service.DoctorService) *DoctorController {
	//     return &DoctorController{Service: s}
	// }
	departmentController := controller.NewDepartmentController(departmentService)
	scheduleController := controller.NewScheduleController(scheduleService)
	appointmentController := controller.NewAppointmentController(appointmentService)

	// Setup routes by passing controllers
	routes.SetupRoutes(
		app,
		patientController,
		doctorController,
		departmentController,
		scheduleController,
		appointmentController,
	)

	err = app.Listen(":3000")
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
