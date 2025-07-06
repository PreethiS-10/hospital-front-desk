package service

import (
	"github.com/PreethiS-10/fiber-crud/internal/models"
	"github.com/PreethiS-10/fiber-crud/internal/repository"
)

type AppointmentService interface {
	GetAll() ([]models.Appointment, error)
	GetByID(id uint) (*models.Appointment, error)
	Create(app *models.Appointment) error
	Update(app *models.Appointment) error
	Delete(id uint) error
}

type appointmentService struct {
	repo repository.AppointmentRepository
}

func NewAppointmentService(r repository.AppointmentRepository) AppointmentService {
	return &appointmentService{repo: r}
}

func (s *appointmentService) GetAll() ([]models.Appointment, error) {
	return s.repo.FindAll()
}

func (s *appointmentService) GetByID(id uint) (*models.Appointment, error) {
	return s.repo.FindByID(id)
}

func (s *appointmentService) Create(app *models.Appointment) error {
	return s.repo.Create(app)
}

func (s *appointmentService) Update(app *models.Appointment) error {
	return s.repo.Update(app)
}

func (s *appointmentService) Delete(id uint) error {
	return s.repo.Delete(id)
}
