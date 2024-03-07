package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"submission_promotion_api/internal/app/handlers"
	"submission_promotion_api/internal/app/services"
)

func HelloServer(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func PromotionRoute(e *echo.Echo, promoService services.PromotionService) {
	e.GET("/", HelloServer)
	e.GET("/promotions", handlers.PSQLGetAllPromotionData(promoService))
	e.GET("/getpromotion/:promotion_id", handlers.PSQLGetPromotionbyPromotionID(promoService))
	e.POST("/createpromotion", handlers.PSQLCreatePromotionData(promoService))
	e.PUT("/updatepromotion/:promotion_id", handlers.PSQLUpdatePromotionbyPromotionID(promoService))
	e.DELETE("/deletepromotion/:promotion_id", handlers.PSQLDeletePromotionbyPromotionID(promoService))
}
