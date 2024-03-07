package repositories

import (
	"gorm.io/gorm"
	"submission_promotion_api/internal/app/models"
)

type PromotionRepository interface {
	CreatePromotion(promo models.Promotion) (models.Promotion, error)
	GetAllPromotions() ([]models.Promotion, error)
	GetPromotionbyPromotionID(promotionID string) (models.Promotion, error)
	UpdatePromotionbyPromotionID(promo models.Promotion) (models.Promotion, error)
	DeletePromotionbyPromotionID(promotionID string) error
}

type PromotionRepositoryImpl struct {
	db *gorm.DB
}

// NewPromotionRepository creates a new instance of PromotionRepository
func NewPromotionRepository(db *gorm.DB) PromotionRepository {
	return &PromotionRepositoryImpl{
		db: db,
	}
}

// CreatePromotion creates a new promotion in the database
func (r *PromotionRepositoryImpl) CreatePromotion(promo models.Promotion) (models.Promotion, error) {
	err := r.db.Create(&promo).Error
	return promo, err
}

// GetAllPromotions returns all promotions recorded in the database
func (r *PromotionRepositoryImpl) GetAllPromotions() ([]models.Promotion, error) {
	var promotions []models.Promotion
	err := r.db.Find(&promotions).Error
	return promotions, err
}

// GetPromotionByPromotionID returns a promotion based on the given promotionID
func (r *PromotionRepositoryImpl) GetPromotionbyPromotionID(PromotionID string) (models.Promotion, error) {
	var promotion models.Promotion
	err := r.db.Where("promotion_id = ?", PromotionID).First(&promotion).Error
	return promotion, err
}

// UpdatePromotionByPromotionID updates a promotion based on the given promotionID
func (r *PromotionRepositoryImpl) UpdatePromotionbyPromotionID(promo models.Promotion) (models.Promotion, error) {
	err := r.db.Save(&promo).Error
	return promo, err
}

// DeletePromotionByPromotionID deletes a promotion based on the given promotionID
func (r *PromotionRepositoryImpl) DeletePromotionbyPromotionID(promotionID string) error {
	err := r.db.Where("promotion_id = ?", promotionID).Delete(&models.Promotion{}).Error
	return err
}
