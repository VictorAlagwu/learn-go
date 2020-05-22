package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/victoralagwu/learn-go/projects/rest-api/api/auth"
	"github.com/victoralagwu/learn-go/projects/rest-api/api/models"
	"github.com/victoralagwu/learn-go/projects/rest-api/api/responses"
	"github.com/victoralagwu/learn-go/projects/rest-api/api/utils/formaterror"
)

func (server *Server) CreatePost(w http.ResponseWriter, r *http.Request) {
	body,err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	post := models.Post{}
	err = json.Unmarshal(body, &post)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	post.Prepare()
	err  = post.Validate()

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	postCreated, err := post.Save(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, postCreated.ID))
	responses.JSON(w, http.StatusCreated, postCreated)
}

//GetPosts :
func (server *Server) GetPosts(w http.ResponseWriter, r *http.Request) {
	post := models.Post{}
	posts, err := post.FindAll(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, posts)
}

//GetPost :
func (server *Server) GetPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	post := models.Post{}
	postDetails, err := post.FindPostByID(server.DB, uint64(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, postDetails)
}

//UpdatePost :
func (server *Server) UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	pid, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//Check if token is valid
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, errors.New("Unauthorized"))
		return
	}

	//Check to see, if post existed
	post := models.Post{}
	err = server.DB.Debug().Model(models.Post{}).Where("id = ?", pid).Take(&post).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Post not found"))
		return
	}

	if uid != post.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Read the data posted
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	postUpdate := models.Post{}
	err = json.Unmarshal(body, &postUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	//Check if the user making request is same as that of the token
	if uid != postUpdate.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	postUpdate.Prepare()
	err = postUpdate.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	postUpdate.ID = post.ID
	postUpdated, err := postUpdate.Update(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	responses.JSON(w, http.StatusOK, postUpdated)
}
//DeletePost :
func (server *Server) DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	post := models.Post{}
	err = server.DB.Debug().Model(models.Post{}).Where("id = ?", pid).Take(&post).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}
	if uid != post.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	_, err = post.Delete(server.DB, pid, uint64(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent, "")
}