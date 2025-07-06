package repository

import (
	"github.com/PreethiS-10/fiber-crud/internal/models"
	"gorm.io/gorm"
)

type AppointmentRepository interface {
	FindAll() ([]models.Appointment, error)
	FindByID(id uint) (*models.Appointment, error)
	Create(appointment *models.Appointment) error
	Update(appointment *models.Appointment) error
	Delete(id uint) error
}

type appointmentRepo struct {
	db *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) AppointmentRepository {
	return &appointmentRepo{db}
}

func (r *appointmentRepo) FindAll() ([]models.Appointment, error) {
	var appointments []models.Appointment
	err := r.db.Preload("Patient").Preload("Doctor").Find(&appointments).Error
	return appointments, err
}

func (r *appointmentRepo) FindByID(id uint) (*models.Appointment, error) {
	var appointment models.Appointment
	err := r.db.Preload("Patient").Preload("Doctor").First(&appointment, id).Error
	return &appointment, err
}

func (r *appointmentRepo) Create(appointment *models.Appointment) error {
	return r.db.Create(appointment).Error
}

func (r *appointmentRepo) Update(appointment *models.Appointment) error {
	return r.db.Save(appointment).Error
}

func (r *appointmentRepo) Delete(id uint) error {
	return r.db.Delete(&models.Appointment{}, id).Error
}
