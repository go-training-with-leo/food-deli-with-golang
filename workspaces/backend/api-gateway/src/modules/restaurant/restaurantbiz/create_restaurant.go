package restaurantbiz

import (
	"api-gateway/src/modules/restaurant/restaurantmodel"
	"context"
	"errors"
	"fmt"
)

type CreateRestaurantStore interface {
	Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	fmt.Println(data.Name)
	if data.Name == "" {
		return errors.New("restaurant name can not be blank")
	}

	err := biz.store.Create(ctx, data)

	return err
}
