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

func (server *Server) CreateDriver(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	driver := models.Driver{}
	err = json.Unmarshal(body, &driver)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	_, err = driver.SaveDriver(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%s", r.Host, r.URL.Path, driver.ID))
	responses.JSON(w, http.StatusCreated, driver)
}

func (server *Server) GetDrivers(w http.ResponseWriter, r *http.Request) {

	driver := models.Driver{}

	drivers, err := driver.FindAllDrivers(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, drivers)
}

func (server *Server) GetDriver(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	did := vars["id"]

	driver := models.Driver{}

	driverReceived, err := driver.FindDriverByID(server.DB, did)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, driverReceived)
}

func (server *Server) UpdateDriver(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	did := vars["id"]

	_, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	driver := models.Driver{}
	err = server.DB.Debug().Model(models.Driver{}).Where("id = ?", did).Take(&driver).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Driver not found"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	driverUpdate := models.Driver{}
	err = json.Unmarshal(body, &driverUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	driverUpdate.ID = driver.ID

	driverUpdated, err := driverUpdate.UpdateDriver(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	responses.JSON(w, http.StatusOK, driverUpdated)
}

func (server *Server) DeleteDriver(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	did := vars["id"]

	_, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	driver := models.Driver{}
	err = server.DB.Debug().Model(models.Driver{}).Where("id = ?", did).Take(&driver).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Driver not found"))
		return
	}

	err = driver.DeleteDriver(server.DB, did)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%s", did))
	responses.JSON(w, http.StatusNoContent, "")
}
