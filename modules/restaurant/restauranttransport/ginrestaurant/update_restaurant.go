package ginrestaurant

import (
	"awesomeProject/common"
	"awesomeProject/component/component"
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
			panic(common.ErrInvalidRequest(err))
		}

		var data restaurantmodel.RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurant(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		c.JSON(401, common.SimpleSuccessResponse(true))
	}

}
