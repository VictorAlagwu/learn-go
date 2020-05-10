package controllers

import (
	"net/http"

	"github.com/victoralagwu/learngo/projects/rest-api/responses"
)

//Home :
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOk, "Welcome to LearnGo API")
}