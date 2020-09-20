package routers

import (
	"net/http"

	template_service "github.com/AlDrac/wallister_test_project/app/web/services"
	"github.com/gorilla/sessions"
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
	//suF, _ := getFlashesSuccess(w, r)
	//if suF != nil {
	//	data["success"] = suF
	//}
	//erF, _ := getFlashesError(w, r)
	//if suF != nil {
	//	data["error"] = erF
	//}

	err := template_service.RenderTemplate(w, "index.tmpl", TemplateData{
		Page: "home",
		Data: data,
	})
	if err != nil {
		return err
	}

	return nil
}

func customersHandler(w http.ResponseWriter, r *http.Request) error {
	var data = make(map[string]interface{})
	//if err := setFlashes(w, r, errorType, "First Error"); err != nil {
	//	return err
	//}
	//if err := setFlashes(w, r, errorType, "First Error"); err != nil {
	//	return err
	//}
	//if err := setFlashes(w, r, successType, "First Success"); err != nil {
	//	return err
	//}
	//http.Redirect(w, r, "/", http.StatusFound)

	err := template_service.RenderTemplate(w, "customers.tmpl", TemplateData{
		Page: "customers",
		Data: data,
	})
	if err != nil {
		return err
	}

	return nil
}

func customerCreateHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func customerPostCreateHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func customerViewHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func customerEditHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func customerPostEditHandler(w http.ResponseWriter, r *http.Request) error {

	return nil
}

func customerDeleteHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}
