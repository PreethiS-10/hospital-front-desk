package service

import (
	"github.com/PreethiS-10/fiber-crud/internal/models"
	"github.com/PreethiS-10/fiber-crud/internal/repository"
)

type DoctorService interface {
	GetAll() ([]models.Doctor, error)
	GetByID(id uint) (*models.Doctor, error)
	Create(doctor *models.Doctor) error
	Update(doctor *models.Doctor) error
	Delete(id uint) error
}

type doctorService struct {
	repo repository.DoctorRepository
}

func NewDoctorService(r repository.DoctorRepository) DoctorService {
	return &doctorService{repo: r}
}

func (s *doctorService) GetAll() ([]models.Doctor, error) {
	return s.repo.FindAll()
}

func (s *doctorService) GetByID(id uint) (*models.Doctor, error) {
	return s.repo.FindByID(id)
}

func (s *doctorService) Create(doctor *models.Doctor) error {
	return s.repo.Create(doctor)
}

func (s *doctorService) Update(doctor *models.Doctor) error {
	return s.repo.Update(doctor)
}

func (s *doctorService) Delete(id uint) error {
	return s.repo.Delete(id)
}
