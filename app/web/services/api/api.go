package serviceApi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/AlDrac/wallister_test_project/app/web/models"
	"net/http"
	"time"
)

type ResStruct struct {
	Error      string      `json:"error"`
	Message    string      `json:"message"`
	StatusCode int         `json:"statusCode"`
	Result     interface{} `json:"result"`
}

var (
	apiUrl     string
	httpClient = &http.Client{Timeout: 10 * time.Second}
)

func InitializeServiceApi(aU string) {
	apiUrl = aU
}

func getFromApiData(url string, method string, data map[string]string) ResStruct {
	jsonStr, _ := json.Marshal(data)
	return getFromApi(url, method, jsonStr)
}

func getFromApiCustomer(url string, method string, data models.Customer) ResStruct {
	jsonStr, _ := json.Marshal(data)
	return getFromApi(url, method, jsonStr)
}

func getFromApi(url string, method string, jsonStr []byte) ResStruct {
	r, err := http.NewRequest(method, apiUrl+url, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
	}
	req, err := httpClient.Do(r)
	if err != nil {
		fmt.Println(err)
	}
	defer req.Body.Close()
	res := ResStruct{}
	_ = json.NewDecoder(req.Body).Decode(&res)
	return res
}

func GetCustomers(vars map[string]string) interface{} {
	data := make(map[string]string)
	data["first_name"] = vars["first_name"]
	data["last_name"] = vars["last_name"]
	result := getFromApiData("/customers", http.MethodGet, data)
	if result.Error != "" {
		return make([]models.Customer, 0)
	}
	return result.Result
}

func GetCustomer(id string) (interface{}, error) {
	data := make(map[string]string)
	result := getFromApiData("/customer/"+id, http.MethodGet, data)
	if result.Error != "" {
		return models.Customer{}, errors.New(result.Message)
	}
	return result.Result, nil
}

func DeleteCustomer(id string) error {
	data := make(map[string]string)
	result := getFromApiData("/customer/delete/"+id, http.MethodDelete, data)
	if result.Error != "" {
		return errors.New(result.Message)
	}
	return nil
}

func UpdateCustomer(customer models.Customer, id string) error {
	result := getFromApiCustomer("/customer/edit/"+id, http.MethodPut, customer)
	if result.Error != "" {
		return errors.New(result.Message)
	}
	return nil
}

func CreateCustomer(customer models.Customer) error {
	result := getFromApiCustomer("/customer/create", http.MethodPost, customer)
	if result.Error != "" {
		return errors.New(result.Message)
	}
	return nil
}
