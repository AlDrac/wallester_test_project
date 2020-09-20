package routers

import (
	"fmt"
	"github.com/AlDrac/wallister_test_project/app/web/models"
	serviceApi "github.com/AlDrac/wallister_test_project/app/web/services/api"
	serviceTemplate "github.com/AlDrac/wallister_test_project/app/web/services/template"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"net/http"
	"strconv"
)

const (
	errorType   = "error"
	successType = "success"
)

var store *sessions.CookieStore

func SetStore(key string) {
	store = sessions.NewCookieStore([]byte(key))
}

func getFlashesError(w http.ResponseWriter, r *http.Request) ([]interface{}, error) {
	session, err := store.Get(r, "flash-session")
	if err != nil {
		return nil, err
	}
	flashes := session.Flashes(errorType)
	if flashes == nil {
		return nil, nil
	}
	if err := session.Save(r, w); err != nil {
		return nil, err
	}
	return flashes, nil
}

func getFlashesSuccess(w http.ResponseWriter, r *http.Request) ([]interface{}, error) {
	session, err := store.Get(r, "flash-session")
	if err != nil {
		return nil, err
	}
	flashes := session.Flashes(successType)
	if flashes == nil {
		return nil, nil
	}
	if err := session.Save(r, w); err != nil {
		return nil, err
	}
	return flashes, nil
}

func setFlashes(w http.ResponseWriter, r *http.Request, flashType string, message string) error {
	session, err := store.Get(r, "flash-session")
	if err != nil {
		return err
	}
	session.AddFlash(message, flashType)
	if err := session.Save(r, w); err != nil {
		return err
	}

	return nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) error {
	var data = make(map[string]interface{})

	err := serviceTemplate.RenderTemplate(w, "index.tmpl", TemplateData{
		Page: "home",
		Data: data,
	})
	if err != nil {
		return err
	}

	return nil
}

func customersHandler(w http.ResponseWriter, r *http.Request) error {
	data := make(map[string]interface{})

	suF, _ := getFlashesSuccess(w, r)
	if suF != nil {
		data[successType] = suF
	}
	erF, _ := getFlashesError(w, r)
	if erF != nil {
		data[errorType] = erF
	}

	search := map[string]string{
		"first_name": r.URL.Query().Get("first_name"),
		"last_name":  r.URL.Query().Get("last_name"),
	}
	data["search"] = search
	customers := serviceApi.GetCustomers(search)
	data["customers"] = customers

	err := serviceTemplate.RenderTemplate(w, "customers.tmpl", TemplateData{
		Page: "customers",
		Data: data,
	})
	if err != nil {
		return err
	}

	return nil
}

func customerCreateHandler(w http.ResponseWriter, r *http.Request) error {
	//data := make(map[string]interface{})
	//err := serviceTemplate.RenderTemplate(w, "customer_create.tmpl", TemplateData{
	//	Page: "customer_create",
	//	Data: data,
	//})
	//if err != nil {
	//	return err
	//}
	return nil
}

func customerPostCreateHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func customerViewHandler(w http.ResponseWriter, r *http.Request) error {
	data := make(map[string]interface{})

	suF, _ := getFlashesSuccess(w, r)
	if suF != nil {
		data[successType] = suF
	}
	erF, _ := getFlashesError(w, r)
	if erF != nil {
		data[errorType] = erF
	}

	customer, err := serviceApi.GetCustomer(mux.Vars(r)["id"])
	if err != nil {
		if err := setFlashes(w, r, errorType, err.Error()); err != nil {
			return err
		}
		http.Redirect(w, r, "/customers", http.StatusFound)
		return nil
	}
	data["customer"] = customer
	if err := serviceTemplate.RenderTemplate(w, "customer_view.tmpl", TemplateData{
		Page: "customer_view",
		Data: data,
	}); err != nil {
		return err
	}

	return nil
}

func customerEditHandler(w http.ResponseWriter, r *http.Request) error {
	data := make(map[string]interface{})

	suF, _ := getFlashesSuccess(w, r)
	if suF != nil {
		data[successType] = suF
	}
	erF, _ := getFlashesError(w, r)
	if erF != nil {
		data[errorType] = erF
	}

	if err := r.ParseForm(); err != nil {
		return err
	}
	if r.FormValue("id") != "" {
		customer := models.Customer{}
		customer.FirstName = r.FormValue("first_name")
		customer.LastName = r.FormValue("last_name")
		customer.BirthDate = r.FormValue("birth_date")
		customer.Gender = r.FormValue("gender")
		customer.Email = r.FormValue("email")
		customer.Address = r.FormValue("address")
		data["formData"] = customer
	}

	customer, err := serviceApi.GetCustomer(mux.Vars(r)["id"])
	if err != nil {
		if err := setFlashes(w, r, errorType, err.Error()); err != nil {
			return err
		}
		http.Redirect(w, r, "/customers", http.StatusFound)
		return nil
	}
	data["customer"] = customer

	if err := serviceTemplate.RenderTemplate(w, "customer_edit.tmpl", TemplateData{
		Page: "customer_edit",
		Data: data,
	}); err != nil {
		return err
	}

	return nil
}

func customerPostEditHandler(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	customer := models.Customer{}
	customer.ID, _ = strconv.Atoi(r.FormValue("id"))
	customer.FirstName = r.FormValue("first_name")
	customer.LastName = r.FormValue("last_name")
	customer.BirthDate = r.FormValue("birth_date")
	customer.Gender = r.FormValue("gender")
	customer.Email = r.FormValue("email")
	customer.Address = r.FormValue("address")

	if err := serviceApi.UpdateCustomer(customer, mux.Vars(r)["id"]); err != nil {
		if err := setFlashes(w, r, errorType, err.Error()); err != nil {
			return err
		}
		if err := customerEditHandler(w, r); err != nil {
			return err
		}
		return nil
	}

	if err := setFlashes(w, r, successType, "The customer has been updated (#"+mux.Vars(r)["id"]+")"); err != nil {
		return err
	}
	http.Redirect(w, r, "/customer/"+mux.Vars(r)["id"], http.StatusFound)

	return nil
}

func customerDeleteHandler(w http.ResponseWriter, r *http.Request) error {
	fmt.Println(mux.Vars(r)["id"])

	if err := serviceApi.DeleteCustomer(mux.Vars(r)["id"]); err != nil {
		if err := setFlashes(w, r, errorType, err.Error()); err != nil {
			return err
		}
		http.Redirect(w, r, "/customer/"+mux.Vars(r)["id"], http.StatusFound)
		return nil
	}

	if err := setFlashes(w, r, successType, "The customer has been deleted (#"+mux.Vars(r)["id"]+")"); err != nil {
		return err
	}
	http.Redirect(w, r, "/customers", http.StatusFound)

	return nil
}
