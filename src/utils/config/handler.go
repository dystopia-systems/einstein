package config

import (
	"bytes"
	"encoding/json"
	"github.com/vectorman1/alaskalog"
	"io/ioutil"
	"net/http"
	"os"
)

var _config *Config

func InitConfig() error {
	alaskalog.Logger.Infoln("Getting configuration from Saruman...")

	client := &http.Client{}
	url := os.Getenv("SARUMAN_URL")
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Api-Key", os.Getenv("SARUMAN_API_KEY"))

	res, err := client.Do(req)

	if err != nil || res.StatusCode != http.StatusOK {
		alaskalog.Logger.Warnln("Request failed or not HTTP 200 OK")
		return err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &_config)

	alaskalog.Logger.Infoln("Configuration loaded.")

	return nil
}

func GetConfig() (*Config, error) {
	if _config != nil {
		return _config, nil
	}

	return nil, nil
}

func SaveConfig(config *Config) {
	file, _ := json.MarshalIndent(&config, "", "	")
	body := bytes.NewReader(file)

	client := &http.Client{}
	url := os.Getenv("SARUMAN_URL")
	req, _ := http.NewRequest("POST", url, body)
	req.Header.Set("Api-Key", os.Getenv("SARUMAN_API_KEY"))

	_, err := client.Do(req)

	if err != nil {
		alaskalog.Logger.Fatalf("Error saving configuration to remote server %v", err)
		return
	}

	_config = config

	alaskalog.Logger.Infoln("Configuration saved.")
}


