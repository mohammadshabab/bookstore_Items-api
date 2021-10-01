package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mohammadshabab/bookstore_items-api/domain/items"
	"github.com/mohammadshabab/bookstore_items-api/services"
	"github.com/mohammadshabab/bookstore_items-api/utils/http_utils"
	"github.com/mohammadshabab/bookstore_utils-go/rest_errors"

	"github.com/mohammadshabab/bookstore_oauth-go/oauth"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct {
}

//To create an item
func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.RespondError(w, err)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, respErr)
		return
	}

	itemRequest.Seller = oauth.GetCallerId(r)

	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		http_utils.RespondError(w, createErr)

		return
	}

	http_utils.RespondJson(w, http.StatusCreated, result)

}

//To retrieve an item
func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
