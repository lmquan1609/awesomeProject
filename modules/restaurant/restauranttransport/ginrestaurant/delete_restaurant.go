package ginrestaurant

import (
	"awesomeProject/common"
	"awesomeProject/component/component"
	"awesomeProject/modules/restaurant/restaurantbiz"
	"awesomeProject/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
	"strconv"
)

func DeleteRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.JSON(401, common.SimpleSuccessResponse(true))
	}

}
