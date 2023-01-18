package restaurantbiz

import (
	"awesomeProject/common"
	"awesomeProject/modules/restaurant/restaurantmodel"
	"context"
	"errors"
)

type GetRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type getRestaurantBiz struct {
	store GetRestaurantStore
}

func NewGetRestaurantBiz(store GetRestaurantStore) *getRestaurantBiz {
	return &getRestaurantBiz{store: store}
}

func (biz *getRestaurantBiz) GetRestaurant(
	ctx context.Context,
	id int,
) (*restaurantmodel.Restaurant, error) {
	result, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err == common.RecordNotFound {
			return nil, err
		}
		return nil, err
	}

	if result.Status == 0 {
		return nil, errors.New("data deleted")
	}
	return result, err
}
