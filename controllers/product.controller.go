package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mhdianrush/go-products-rest-api/config"
	"github.com/mhdianrush/go-products-rest-api/entities"
	"github.com/mhdianrush/go-products-rest-api/helper"
	"gorm.io/gorm"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var products []entities.Product

	if err := config.DB.Find(&products).Error; err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	helper.ResponseJSON(w, http.StatusOK, products)
}

func Find(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var product entities.Product

	if err := config.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.ResponseError(w, http.StatusNotFound, "product not found")
			return
		default:
			helper.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	helper.ResponseJSON(w, http.StatusOK, product)
}

func Create(w http.ResponseWriter, r *http.Request) {

}

func Update(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
