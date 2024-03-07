package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"submission_promotion_api/internal/app/models"
	"submission_promotion_api/internal/app/services"
)

func PSQLCreatePromotionData(PromoService services.PromotionService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var promo models.Promotion
		if err := c.Bind(&promo); err != nil {
			return err
		}
		createdPromo, err := PromoService.CreatePromotion(promo)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, createdPromo)
	}
}

func PSQLGetAllPromotionData(PromoService services.PromotionService) echo.HandlerFunc {
	return func(c echo.Context) error {
		promos, err := PromoService.GetAllPromotions()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, promos)
	}
}

func PSQLGetPromotionbyPromotionID(PromoService services.PromotionService) echo.HandlerFunc {
	return func(c echo.Context) error {
		promotionID := c.Param("promotion_id")
		promo, err := PromoService.GetPromotionbyPromotionID(promotionID)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, promo)
	}
}

func PSQLUpdatePromotionbyPromotionID(PromoService services.PromotionService) echo.HandlerFunc {
	return func(c echo.Context) error {
		promotionID := c.Param("promotion_id")
		var promo models.Promotion
		if err := c.Bind(&promo); err != nil {
			return err
		}
		promo.PromotionID = promotionID
		updatedPromo, err := PromoService.UpdatePromotionbyPromotionID(promo)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, updatedPromo)
	}
}

func PSQLDeletePromotionbyPromotionID(PromoService services.PromotionService) echo.HandlerFunc {
	return func(c echo.Context) error {
		promotionID := c.Param("promotion_id")
		if err := PromoService.DeletePromotionbyPromotionID(promotionID); err != nil {
			return err
		}
		return c.NoContent(http.StatusNoContent)
	}
}
