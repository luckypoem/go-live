package restfulapi

import (
	"go-live/models"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// App Restful API
func CreateAppHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	appname := ps.ByName("appname")
	liveon := r.FormValue("liveon")
	if liveon == "" {
		liveon = "on"
	}

	err := models.CreateApp(&models.App{
		Appname: appname,
		Liveon:  liveon,
	})

	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendResponse(w, &Response{
		Code:    http.StatusOK,
		Message: "Successfully created this app.",
	})
}

func ListAppsHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	apps, err := models.GetAllApps()

	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendResponse(w, &AppsResponse{
		Code:    http.StatusOK,
		Data:    apps,
		Message: "Successfully acquired all applications.",
	})
}

func GetAppByIdHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	appid := ps.ByName("appid")
	if appid == "" {
		SendErrorResponse(w, http.StatusBadRequest, "Appid is not be null.")
		return
	}

	id, err := strconv.Atoi(appid)

	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	app, err := models.GetAppById(id)

	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendResponse(w, &AppResponse{
		Code:    http.StatusOK,
		Data:    app,
		Message: "Successfully obtained the corresponding application.",
	})
}

func DeleteAppByIdHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	appid := ps.ByName("appid")
	if appid == "" {
		SendErrorResponse(w, http.StatusBadRequest, "Appid is not be null.")
		return
	}

	id, err := strconv.Atoi(appid)

	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = models.DeleteApp(&models.App{Id: id})

	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendResponse(w, &Response{
		Code:    http.StatusOK,
		Message: "Successfully deleted this app.",
	})
}

// Live Restful API
func CreateLiveHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func ListLivesHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func GetLiveByIdHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func UpdateLiveTokenByIdHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func DeleteLiveByIdHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
