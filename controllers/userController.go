package controllers

import (
	"encoding/json"
	"net/http"
	"spotestapi/common"
	"spotestapi/data"
	"spotestapi/models"
)

//Register Handler for HTTP Post - "/users/register"
//Add a new User document
func Register(w http.ResponseWriter, r *http.Request) {
	var dataResource UserResource
	//Decode the incoming User json

	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid User data",
			500,
		)
		return
	}
	user := &dataResource.Data

	datab := GetDatabaseWithContext()
	col := datab.DbCollection("users")
	repo := &data.UserRepository{C: col}

	//insert User document
	repo.CreateUser(user)

	//Clean-up the hashpassword to eliminate it from response
	user.HashPassword = nil

	if j, err := json.Marshal(UserResource{Data: *user}); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}

}

//Login Handler for HTTP Post request - "user/login"
//Login user
func Login(w http.ResponseWriter, r *http.Request) {
	var dataResource LoginResource
	var token string

	//Decode incoming Login json
	err := json.NewDecoder(r.Body).Decode(&dataResource)

	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Login Data",
			500,
		)
		return
	}

	loginModel := dataResource.Data
	loginUser := models.User{
		Email:    loginModel.Email,
		Password: loginModel.Password,
	}

	datab := GetDatabaseWithContext()
	col := datab.DbCollection("users")
	repo := &data.UserRepository{C: col}

	//Authenticate the login user
	user, err := repo.Login(loginUser)

	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Login Credentials",
			401,
		)
	}

	token, err = common.GenerateJWT(user.Email, "member")
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Eror while generating the access token",
			500,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	//Clean-up hashpassword to eliminate it from response JSON
	user.HashPassword = nil

	authUser := AuthUserModel{
		User:  user,
		Token: token,
	}
	j, err := json.Marshal(AuthUserResource{Data: authUser})

	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(j)

}
