package services

import (
	"github.com/mohammadshabab/bookstore_items-api/src/domain/items"
	"github.com/mohammadshabab/bookstore_items-api/src/domain/queries"
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
	Search(queries.EsQuery) ([]items.Item, rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Get(id string) (*items.Item, rest_errors.RestErr) {
	item := items.Item{Id: id}

	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Search(query queries.EsQuery) ([]items.Item, rest_errors.RestErr) {
	dao := items.Item{}
	return dao.Search(query)
}
