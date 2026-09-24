package models

import "time"

type Job struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"size:150;not null;index" json:"title"`
	Description string    `gorm:"type:text;not null" json:"description"`
	Company     string    `gorm:"size:150;not null;index" json:"company"`
	Location    string    `gorm:"size:150;not null;index" json:"location"`
	Salary      float64   `gorm:"not null;index" json:"salary"`
	JobType     string    `gorm:"size:30;not null;default:'full-time'" json:"job_type"`
	PostedBy    uint      `gorm:"not null;index" json:"posted_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Recruiter   User      `gorm:"foreignKey:PostedBy" json:"recruiter"`
}
