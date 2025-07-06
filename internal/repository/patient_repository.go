package repository

import (
	"github.com/PreethiS-10/fiber-crud/internal/models"
	"gorm.io/gorm"
)

type PatientRepository interface {
	FindAll() ([]models.Patient, error)
	FindByID(id uint) (*models.Patient, error)
	Create(patient *models.Patient) error
	Update(patient *models.Patient) error
	Delete(id uint) error
}

type patientRepo struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) PatientRepository {
	return &patientRepo{db}
}

func (r *patientRepo) FindAll() ([]models.Patient, error) {
	var patients []models.Patient
	err := r.db.Find(&patients).Error
	return patients, err
}

func (r *patientRepo) FindByID(id uint) (*models.Patient, error) {
	var patient models.Patient
	err := r.db.First(&patient, id).Error
	return &patient, err
}

func (r *patientRepo) Create(patient *models.Patient) error {
	return r.db.Create(patient).Error
}

func (r *patientRepo) Update(patient *models.Patient) error {
	return r.db.Save(patient).Error
}

func (r *patientRepo) Delete(id uint) error {
	return r.db.Delete(&models.Patient{}, id).Error
}
