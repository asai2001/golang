package services

import (
	"submission_promotion_api/internal/app/models"
	"submission_promotion_api/internal/app/repositories"
)

// PromotionService provides promotion-related services
type PromotionService interface {
	CreatePromotion(promo models.Promotion) (models.Promotion, error)
	GetAllPromotions() ([]models.Promotion, error)
	GetPromotionbyPromotionID(promotionID string) (models.Promotion, error)
	UpdatePromotionbyPromotionID(promo models.Promotion) (models.Promotion, error)
	DeletePromotionbyPromotionID(promotionID string) error
}

type PromotionServiceImpl struct {
	PromotionRepo repositories.PromotionRepository
}

// NewPromotionService creates a new instance of PromotionService
func NewPromotionService(PromotionRepo repositories.PromotionRepository) *PromotionServiceImpl {
	return &PromotionServiceImpl{
		PromotionRepo: PromotionRepo,
	}
}

// CreatePromotion creates a new promotion
func (s *PromotionServiceImpl) CreatePromotion(promo models.Promotion) (models.Promotion, error) {
	createdPromo, err := s.PromotionRepo.CreatePromotion(promo)
	if err != nil {
		return models.Promotion{}, err
	}
	return createdPromo, nil
}

// GetPromotionByPromotionID returns a promotion based on the given promotionID
func (s *PromotionServiceImpl) GetPromotionbyPromotionID(promotionID string) (models.Promotion, error) {
	promotion, err := s.PromotionRepo.GetPromotionbyPromotionID(promotionID)
	if err != nil {
		return models.Promotion{}, err
	}
	return promotion, nil
}

// UpdatePromotionbyPromotionID updates a promotion based on the given promotionID
func (s *PromotionServiceImpl) UpdatePromotionbyPromotionID(promo models.Promotion) (models.Promotion, error) {
	updatedPromo, err := s.PromotionRepo.UpdatePromotionbyPromotionID(promo)
	if err != nil {
		return models.Promotion{}, err
	}
	return updatedPromo, nil
}

// DeletePromotionByPromotionID deletes a promotion based on the given promotionID
func (s *PromotionServiceImpl) DeletePromotionbyPromotionID(promotionID string) error {
	err := s.PromotionRepo.DeletePromotionbyPromotionID(promotionID)
	if err != nil {
		return err
	}
	return nil
}

// GetAllPromotions retrieves all promotions from the database
func (s *PromotionServiceImpl) GetAllPromotions() ([]models.Promotion, error) {
	promotions, err := s.PromotionRepo.GetAllPromotions()
	if err != nil {
		return nil, err
	}
	return promotions, nil
}
