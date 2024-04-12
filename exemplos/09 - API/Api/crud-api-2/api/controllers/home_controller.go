package controllers

import (
	"net/http"

	"github.com/luizgambali/crud-api-2/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
