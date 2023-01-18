package ginrestaurant

import (
	"awesomeProject/common"
	"awesomeProject/component"
	"awesomeProject/modules/restaurant/restaurantbiz"
	"awesomeProject/modules/restaurant/restaurantmodel"
	"awesomeProject/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
	"strconv"
)

func UpdateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		var data restaurantmodel.RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurant(c.Request.Context(), id, &data); err != nil {
			return
		}

		c.JSON(401, common.SimpleSuccessResponse(true))
	}

}
