package common

import "go.uber.org/zap"

//StartUp bootstraps the initial configurations of app
func StartUp(log *zap.Logger) error {

	//Initialize AppConfig variable
	err := initConfig(log)
	if err != nil {
		return err
	}
	//initialize private and public keys for JWT authentication
	err = initKeyog(log)
	if err != nil {
		return err
	}

	return nil

}
