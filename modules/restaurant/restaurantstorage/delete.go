package restaurantstorage

import (
	"awesomeProject/common"
	"awesomeProject/modules/restaurant/restaurantmodel"
	"context"
)

func (s *sqlStore) SoftDeleteData(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
