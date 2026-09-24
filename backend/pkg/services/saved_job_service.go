package services

import (
	"errors"

	"github.com/pranav-patidar9354/job-portal-v2/backend/pkg/models"
	"github.com/pranav-patidar9354/job-portal-v2/backend/pkg/repositories"
)

type SavedJobService struct {
	Saved *repositories.SavedJobRepository
	Jobs  *repositories.JobRepository
}

func (s *SavedJobService) Toggle(jobID, candidateID uint) (bool, error) {
	if _, err := s.Jobs.FindByID(jobID); err != nil {
		return false, errors.New("job not found")
	}

	saved, err := s.Saved.Find(jobID, candidateID)
	if err == nil {
		if err := s.Saved.Delete(saved); err != nil {
			return false, err
		}
		return false, nil
	}

	if err := s.Saved.Create(&models.SavedJob{
		JobID: jobID, CandidateID: candidateID,
	}); err != nil {
		return false, err
	}
	return true, nil
}

func (s *SavedJobService) List(id uint) ([]models.SavedJob, error) {
	return s.Saved.List(id)
}
