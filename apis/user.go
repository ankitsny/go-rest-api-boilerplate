package apis

import (
	"encoding/json"
	"fmt"
	"goapi/app"
	"goapi/models"
	"net/http"

	"github.com/gorilla/mux"
)

type userService interface {
	Get(rs app.RequestScope, email string) (*models.User, error)
	Count(rs app.RequestScope) (int, error)
	Create(rs app.RequestScope, model *models.User) error
	Update(rs app.RequestScope, email string, model *models.User) (*models.User, error)
	Delete(rs app.RequestScope, email string) error
}

type userResource struct {
	service userService
}

// ServeUserResource sets up the routing of user endpoints and the corresponding handlers.
func ServeUserResource(rg *mux.Router, service userService) {
	r := &userResource{service}
	rg.HandleFunc("/users/{email}", r.get).Methods("GET")
	rg.HandleFunc("/users", r.create).Methods("POST")
	rg.HandleFunc("/users/{email}", r.update).Methods("PUT")
	rg.HandleFunc("/users/{email}", r.delete).Methods("DELETE")
}

func (ur *userResource) get(w http.ResponseWriter, r *http.Request) {
	rs := app.GetRequestScope(r)
	email := rs.GetParams()["email"]
	response, err := ur.service.Get(rs, email)
	if err != nil {
		fmt.Fprintf(w, "Err")
		return
	}

	uB, _ := json.Marshal(response)

	w.Write(uB)
}

func (ur *userResource) create(w http.ResponseWriter, r *http.Request) {

	rs := app.GetRequestScope(r)
	var x *models.User
	json.Unmarshal(rs.GetBody(), &x)
	err := ur.service.Create(rs, x)

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	fmt.Fprintf(w, "Done")
}

func (ur *userResource) update(w http.ResponseWriter, r *http.Request) {
	rs := app.GetRequestScope(r)
	var x *models.User
	json.Unmarshal(rs.GetBody(), &x)
	u, err := ur.service.Update(rs, rs.GetParams()["email"], x)
	if err != nil {
		fmt.Println("err")
		fmt.Fprintf(w, "Failed to update")
		return
	}
	fmt.Printf("%+v", u)
	fmt.Fprintf(w, "update")
}

func (ur *userResource) delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete")
	rs := app.GetRequestScope(r)
	if err := ur.service.Delete(rs, rs.GetParams()["email"]); err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, "delete")
}
