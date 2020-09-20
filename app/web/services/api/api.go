package serviceApi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/AlDrac/wallister_test_project/app/web/models"
	"net/http"
	"time"
)

var (
	apiUrl     string
	httpClient = &http.Client{Timeout: 10 * time.Second}
)

func InitializeServiceApi(aU string) {
	apiUrl = aU
}

func getFromApi(url string, method string, target interface{}, data map[string]string) interface{} {
	jsonStr, _ := json.Marshal(data)
	r, err := http.NewRequest(method, apiUrl+url, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
	}
	req, err := httpClient.Do(r)
	if err != nil {
		fmt.Println(err)
	}
	defer req.Body.Close()
	_ = json.NewDecoder(req.Body).Decode(&target)
	return target
}

func GetCustomers(vars map[string]string) interface{} {
	data := make(map[string]string)
	data["first_name"] = vars["first_name"]
	data["last_name"] = vars["last_name"]
	return getFromApi("/customers", http.MethodGet, make([]models.Customer, 0), data)
}
