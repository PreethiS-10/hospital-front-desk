package repository

import (
	"github.com/PreethiS-10/fiber-crud/internal/models"
	"gorm.io/gorm"
)

type DoctorRepository interface {
	FindAll() ([]models.Doctor, error)
	FindByID(id uint) (*models.Doctor, error)
	Create(doctor *models.Doctor) error
	Update(doctor *models.Doctor) error
	Delete(id uint) error
}

type doctorRepo struct {
	db *gorm.DB
}

func NewDoctorRepository(db *gorm.DB) DoctorRepository {
	return &doctorRepo{db}
}

func (r *doctorRepo) FindAll() ([]models.Doctor, error) {
	var doctors []models.Doctor
	err := r.db.Preload("Department").Find(&doctors).Error
	return doctors, err
}

func (r *doctorRepo) FindByID(id uint) (*models.Doctor, error) {
	var doctor models.Doctor
	err := r.db.Preload("Department").First(&doctor, id).Error
	return &doctor, err
}

func (r *doctorRepo) Create(doctor *models.Doctor) error {
	return r.db.Create(doctor).Error
}

func (r *doctorRepo) Update(doctor *models.Doctor) error {
	return r.db.Save(doctor).Error
}

func (r *doctorRepo) Delete(id uint) error {
	return r.db.Delete(&models.Doctor{}, id).Error
}
