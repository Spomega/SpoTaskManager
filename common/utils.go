package common

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"go.uber.org/zap"
)

type key string
type configLiteral struct {
	Host, Username, Password, Database key
}

type configuration struct {
	Server, MongoDBHost, MongoDBUser, MongoDBPwd, Database string
}

type (
	appError struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HttpStatus int    `json:"status"`
	}
	errorResource struct {
		Data appError `json:"data"`
	}
)

//AppConfig holds the configuration values from config.json file

var AppConfig configuration
var AppConfigLiteral configLiteral

//initialize AppConfig

func initConfig(log *zap.Logger) error {
	return loadAppConfig(log)
}

//Reads config.json and decode into Appconfig

func loadAppConfig(log *zap.Logger) error {

	file, err := os.Open("common/config.json")
	defer file.Close()

	if err != nil {
		log.Warn("[loadConfig]:An error occurred while parsing config file")
		return err
	}

	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)

	AppConfigLiteral = configLiteral{
		Host:     key("host"),
		Username: key("username"),
		Password: key("password"),
		Database: key("database"),
	}

	if err != nil {
		log.Warn("[loadConfig]:An error occurred while decoding config file")
		return err
	}

	return nil
}

//
func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int) {

	errObj := appError{
		Error:      handlerError.Error(),
		Message:    message,
		HttpStatus: code,
	}

	log.Printf("[AppError]: %s\n", handlerError)
	w.Header().Set("Content-Type", "application/json: charset=utf-8")
	w.WriteHeader(code)

	if j, err := json.Marshal(errorResource{Data: errObj}); err == nil {
		w.Write(j)
	}
}
