package controller

import (
	"encoding/json"
	"go-restful-api/helper"
	"go-restful-api/model/web"
	"go-restful-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func (contoller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// TODO: create helper function to read from request body. looping code
	decoder := json.NewDecoder(request.Body)
	categoryCreateRequest := web.CategoryCreateRequest{}
	err := decoder.Decode(&categoryCreateRequest)
	helper.PanicIfError(err)

	categoryResponse := contoller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponseStruct := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	// TODO: create helper function to write json body. looping code
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponseStruct)
	helper.PanicIfError(err)
}

func (contoller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	err := decoder.Decode(&categoryUpdateRequest)
	helper.PanicIfError(err)

	categoryId := params.ByName("categoryId")
	id, _ := strconv.Atoi(categoryId)
	categoryUpdateRequest.Id = id

	categoryResponse := contoller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	webResponseStruct := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponseStruct)
	helper.PanicIfError(err)
}

func (contoller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, _ := strconv.Atoi(categoryId)

	contoller.CategoryService.Delete(request.Context(), id)
	webResponseStruct := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponseStruct)
	helper.PanicIfError(err)
}

func (contoller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, _ := strconv.Atoi(categoryId)

	categoryResponse := contoller.CategoryService.FindById(request.Context(), id)
	webResponseStruct := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponseStruct)
	helper.PanicIfError(err)
}

func (contoller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponse := contoller.CategoryService.FindAll(request.Context())
	webResponseStruct := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponseStruct)
	helper.PanicIfError(err)
}
