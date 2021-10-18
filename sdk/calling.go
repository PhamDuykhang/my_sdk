package sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	apiVersion = "v1"
)

func GetProducts() []Product {
	client := http.DefaultClient

	res, err := client.Get(fmt.Sprintf("http://localhost:8080/api/%s/products", apiVersion))

	if err != nil {
		fmt.Printf("can't call %v", err)
		return nil
	}
	if res.StatusCode != 200 {
		fmt.Printf("can't call %v", res.StatusCode)
		return nil
	}

	data, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	var products []Product
	_ = json.Unmarshal(data, &products)

	return products
}
