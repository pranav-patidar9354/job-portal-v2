package models

import "time"

type SavedJob struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	JobID       uint      `gorm:"not null;uniqueIndex:idx_saved_job_user" json:"job_id"`
	CandidateID uint      `gorm:"not null;uniqueIndex:idx_saved_job_user" json:"candidate_id"`
	CreatedAt   time.Time `json:"created_at"`
	Job         Job       `gorm:"foreignKey:JobID" json:"job"`
}
