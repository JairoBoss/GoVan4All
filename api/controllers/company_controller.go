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

func (server *Server) CreateCompany(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	company := models.Company{}
	err = json.Unmarshal(body, &company)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	_, err = company.SaveCompany(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%s", r.Host, r.URL.Path, company.ID))
	responses.JSON(w, http.StatusCreated, company)
}

func (server *Server) GetCompanies(w http.ResponseWriter, r *http.Request) {

	company := models.Company{}

	companies, err := company.FindAllCompanies(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, companies)
}

func (server *Server) GetCompany(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	cid := vars["id"]

	company := models.Company{}

	companyReceived, err := company.FindCompanyByID(server.DB, cid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, companyReceived)
}

func (server *Server) UpdateCompany(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	cid := vars["id"]

	_, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	company := models.Company{}
	err = server.DB.Debug().Model(models.Company{}).Where("id = ?", cid).Take(&company).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Company not found"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	companyUpdate := models.Company{}
	err = json.Unmarshal(body, &companyUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	companyUpdate.ID = company.ID

	companyUpdated, err := companyUpdate.UpdateCompany(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	responses.JSON(w, http.StatusOK, companyUpdated)
}

func (server *Server) DeleteCompany(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	cid := vars["id"]

	_, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	company := models.Company{}
	err = server.DB.Debug().Model(models.Company{}).Where("id = ?", cid).Take(&company).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Company not found"))
		return
	}

	err = company.DeleteCompany(server.DB, cid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%s", cid))
	responses.JSON(w, http.StatusNoContent, "")
}
