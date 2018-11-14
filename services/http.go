package services

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func PostFormData(url string,values url.Values) (res string,err error){
	resp, err := http.PostForm(url, values)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}
