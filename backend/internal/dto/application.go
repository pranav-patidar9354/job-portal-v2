package dto

type ApplyRequest struct {
	CoverLetter string `json:"cover_letter" binding:"max=5000"`
}

type UpdateApplicationStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=applied reviewing shortlisted accepted rejected"`
}
