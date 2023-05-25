package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hcastellanos-dev/fullstack/api/auth"
	"github.com/hcastellanos-dev/fullstack/api/models"
	"github.com/hcastellanos-dev/fullstack/api/responses"
	"github.com/hcastellanos-dev/fullstack/api/utils/formaterror"
)

func (server *Server) CreateRoute(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	route := models.Route{}
	err = json.Unmarshal(body, &route)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	_, err = route.SaveRoute(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%s", r.Host, r.URL.Path, route.ID))
	responses.JSON(w, http.StatusCreated, route)
}

func (server *Server) GetRoutes(w http.ResponseWriter, r *http.Request) {

	route := models.Route{}

	routes, err := route.FindAllRoutes(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, routes)
}

func (server *Server) GetRoute(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	rid := vars["id"]

	route := models.Route{}

	routeReceived, err := route.FindRouteByID(server.DB, rid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, routeReceived)
}

func (server *Server) UpdateRoute(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	rid := vars["id"]

	_, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	route := models.Route{}
	err = server.DB.Debug().Model(models.Route{}).Where("id = ?", rid).Take(&route).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Route not found"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	routeUpdate := models.Route{}
	err = json.Unmarshal(body, &routeUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	routeUpdate.ID = route.ID

	routeUpdated, err := routeUpdate.UpdateRoute(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	responses.JSON(w, http.StatusOK, routeUpdated)
}

func (server *Server) DeleteRoute(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	rid := vars["id"]

	_, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	route := models.Route{}
	err = server.DB.Debug().Model(models.Route{}).Where("id = ?", rid).Take(&route).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Route not found"))
		return
	}

	err = route.DeleteRoute(server.DB, rid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%s", rid))
	responses.JSON(w, http.StatusNoContent, "")
}
