package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	maindata "github.com/trangnnp-ts/assignment00/Model"
)

type UserRequest struct {
	Full  string `json:"full"`
	Short string `json:"short"`
}

func InputData(w http.ResponseWriter, r *http.Request) {
	var a UserRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	result := maindata.Add(a.Full, a.Short)
	respondWithJson(w, http.StatusOK, result)
}

func GetOne(w http.ResponseWriter, r *http.Request) {
	alias := chi.URLParam(r, "alias")
	result := maindata.GetOne(alias)
	respondWithJson(w, http.StatusOK, result)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	result := maindata.GetAll()
	respondWithJson(w, http.StatusOK, result)

}

func Update(w http.ResponseWriter, r *http.Request) {
	alias := chi.URLParam(r, "alias")

	var a maindata.Dataa
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	result := maindata.Update(alias, a)
	respondWithJson(w, http.StatusOK, result)
}

func Remove(w http.ResponseWriter, r *http.Request) {
	alias := chi.URLParam(r, "alias")
	result := maindata.Remove(alias)
	respondWithJson(w, http.StatusOK, result)
}

func respondWithError(res http.ResponseWriter, code int, mes string) {
	respondWithJson(res, code, map[string]string{"error": mes})
}

func respondWithJson(res http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(code)
	res.Write(response)
}
