package controllers

import "spotestapi/models"

type (
	//UserResource For Post - /user/register
	UserResource struct {
		Data models.User `json:"data"`
	}

	//LoginResource For Post - /user/login
	LoginResource struct {
		Data LoginModel `json:"data"`
	}

	//AuthUserResource Response for authorized user Post -/user/login
	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}

	//LoginModel for authentication
	LoginModel struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	//AuthUserModel for authorized user access token
	AuthUserModel struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}
)
