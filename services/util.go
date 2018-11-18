package services

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func PostFormData(url string,values url.Values) (res []byte,err error){
	resp, err := http.PostForm(url, values)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func DownloadFile(filepath string, url string) error {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreatePath(path string) (error) {
	if exist, err := PathExists(path); err != nil {
		return err
	}else if !exist {
		if err := os.Mkdir(path,os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}