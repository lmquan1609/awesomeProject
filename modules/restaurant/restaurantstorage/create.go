package restaurantstorage

import (
	"awesomeProject/modules/restaurant/restaurantmodel"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
