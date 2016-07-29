package main

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"errors"
	"strconv"
	"path/filepath"
	"os"
)

func RequestJSON(url string) (map[string]interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Unknown server error (" + strconv.Itoa(resp.StatusCode) + ").")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func RequestAPI(url string) (map[string]interface{}, error) {
	data, err := RequestJSON(url)
	if err != nil {
		return nil, err
	}
	if !CheckRequestStatus(data) {
		return nil, errors.New("Unknown server error")
	}
	return data, nil
}

func CheckRequestStatus(data map[string]interface{}) bool {
	return data["status"] == "success"
}

func RequestToPath(url, path string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	parent := filepath.Join(path, "..")
	err = os.MkdirAll(parent, 0777)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, data, 0777)
	if err != nil {
		return err
	}
	return nil
}
