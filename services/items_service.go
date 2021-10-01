package services

import (
	"net/http"

	"github.com/mohammadshabab/bookstore_items-api/domain/items"
	"github.com/mohammadshabab/bookstore_utils-go/rest_errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{} //this means itemService is of time itemServiceInterface and the value
	// of this itemsservice it a pointer to itemsService being a struct
	//And we are doing this for safety means if we will forget or not impletement
	//a method available in itemsServiceInterface it will through error
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("implement me", http.StatusNotImplemented, "not implemented", nil)
}

func (s *itemsService) Get(st string) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("implement me", http.StatusNotImplemented, "not implemented", nil)
}
