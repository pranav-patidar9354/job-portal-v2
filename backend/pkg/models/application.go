package models

import "time"

type Application struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	JobID       uint      `gorm:"not null;index;uniqueIndex:idx_job_candidate" json:"job_id"`
	CandidateID uint      `gorm:"not null;index;uniqueIndex:idx_job_candidate" json:"candidate_id"`
	Status      string    `gorm:"size:30;not null;default:'applied';index" json:"status"`
	CoverLetter string    `gorm:"type:text" json:"cover_letter"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Job         Job       `gorm:"foreignKey:JobID" json:"job"`
	Candidate   User      `gorm:"foreignKey:CandidateID" json:"candidate"`
}
