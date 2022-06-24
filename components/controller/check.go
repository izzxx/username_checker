package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/izzxx/username_checker/components/service"
	"github.com/izzxx/username_checker/components/utility"
)

func UsernameChecker(w http.ResponseWriter, r *http.Request) {
	// get username from URLParam
	username := chi.URLParam(r, "username")

	resp, err := service.UsernameChecker(r, username)
	if err != nil {
		err = utility.ResponseApi(w, http.StatusNotFound, err.Error(), resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	err = utility.ResponseApi(w, http.StatusOK, "found", resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
