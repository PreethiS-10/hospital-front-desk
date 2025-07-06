package service

import (
	"github.com/PreethiS-10/fiber-crud/internal/models"
	"github.com/PreethiS-10/fiber-crud/internal/repository"
)

type PatientService interface {
	GetAll() ([]models.Patient, error)
	GetByID(id uint) (*models.Patient, error)
	Create(patient *models.Patient) error
	Update(patient *models.Patient) error
	Delete(id uint) error
}

type patientService struct {
	repo repository.PatientRepository
}

func NewPatientService(r repository.PatientRepository) PatientService {
	return &patientService{repo: r}
}

func (s *patientService) GetAll() ([]models.Patient, error) {
	return s.repo.FindAll()
}

func (s *patientService) GetByID(id uint) (*models.Patient, error) {
	return s.repo.FindByID(id)
}

func (s *patientService) Create(patient *models.Patient) error {
	return s.repo.Create(patient)
}

func (s *patientService) Update(patient *models.Patient) error {
	return s.repo.Update(patient)
}

func (s *patientService) Delete(id uint) error {
	return s.repo.Delete(id)
}
