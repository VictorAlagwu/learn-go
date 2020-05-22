package controllers

import (
	"net/http"

	"github.com/victoralagwu/learn-go/projects/rest-api/api/responses"
)

//Home :
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to LearnGo API")
}