package models

import "time"

type Patient struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`    
	Email     string    `gorm:"not null;unique" json:"email"` 
	Age 	 int       `gorm:"not null" json:"age"`
	Phone     string    `gorm:"not null" json:"phone"`    
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Doctor struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	Name           string     `gorm:"not null" json:"name"`       
	Email          string     `gorm:"not null;unique" json:"email"`  
	Specialization string     `gorm:"not null" json:"specialization"` 
	DepartmentID   uint       `gorm:"not null" json:"department_id"`  
	Department     Department `gorm:"foreignKey:DepartmentID" json:"department"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

type Department struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null;unique" json:"name"` 
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Schedule struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Day       string    `gorm:"not null" json:"day"`        
	StartTime time.Time `gorm:"not null" json:"start_time"` 
	EndTime   time.Time `gorm:"not null" json:"end_time"`   
	DoctorID  uint      `gorm:"not null" json:"doctor_id"`  
	Doctor    Doctor    `gorm:"foreignKey:DoctorID" json:"doctor"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Appointment struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	AppointmentDate time.Time `gorm:"not null" json:"appointment_date"` 
	AppointmentTime time.Time `gorm:"not null" json:"appointment_time"` 
	Status          string    `gorm:"not null" json:"status"`            
	PatientID       uint      `gorm:"not null" json:"patient_id"`        
	Patient         Patient   `gorm:"foreignKey:PatientID" json:"patient"`
	DoctorID        uint      `gorm:"not null" json:"doctor_id"`         
	Doctor          Doctor    `gorm:"foreignKey:DoctorID" json:"doctor"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
