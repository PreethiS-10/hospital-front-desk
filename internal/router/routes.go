package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/PreethiS-10/fiber-crud/internal/controller"
)

func SetupRoutes(
	app *fiber.App,
	patientCtrl *controller.PatientController,
	doctorCtrl *controller.DoctorController,
	departmentCtrl *controller.DepartmentController,
	scheduleCtrl *controller.ScheduleController,
	appointmentCtrl *controller.AppointmentController,
) {
	// Patient routes
	patient := app.Group("/patients")
	patient.Get("/", patientCtrl.GetPatients)
	patient.Get("/:id", patientCtrl.GetPatient)
	patient.Post("/", patientCtrl.CreatePatient)
	patient.Put("/:id", patientCtrl.UpdatePatient)
	patient.Delete("/:id", patientCtrl.DeletePatient)

	// Doctor routes
	doctor := app.Group("/doctors")
	doctor.Get("/", doctorCtrl.GetDoctors)
	doctor.Get("/:id", doctorCtrl.GetDoctor)
	doctor.Post("/", doctorCtrl.CreateDoctor)
	doctor.Put("/:id", doctorCtrl.UpdateDoctor)
	doctor.Delete("/:id", doctorCtrl.DeleteDoctor)

	// Department routes
	department := app.Group("/departments")
	department.Get("/", departmentCtrl.GetDepartments)
	department.Get("/:id", departmentCtrl.GetDepartment)
	department.Post("/", departmentCtrl.CreateDepartment)
	department.Put("/:id", departmentCtrl.UpdateDepartment)
	department.Delete("/:id", departmentCtrl.DeleteDepartment)

	// Schedule routes
	schedule := app.Group("/schedules")
	schedule.Get("/", scheduleCtrl.GetSchedules)
	schedule.Get("/:id", scheduleCtrl.GetSchedule)
	schedule.Post("/", scheduleCtrl.CreateSchedule)
	schedule.Put("/:id", scheduleCtrl.UpdateSchedule)
	schedule.Delete("/:id", scheduleCtrl.DeleteSchedule)

	// Appointment routes
	appointment := app.Group("/appointments")
	appointment.Get("/", appointmentCtrl.GetAppointments)
	appointment.Get("/:id", appointmentCtrl.GetAppointment)
	appointment.Post("/", appointmentCtrl.CreateAppointment)
	appointment.Put("/:id", appointmentCtrl.UpdateAppointment)
	appointment.Delete("/:id", appointmentCtrl.DeleteAppointment)
}
