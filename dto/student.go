package dto

import "github.com/flowers-bloom/gin_web_scaffold/model"

type LoginDto struct {
	Email string `json:"email" binding:"required,max=20"`
	Pwd   string `json:"pwd" binding:"required,max=32"`
}

type LoginResult struct {
	Token string `json:"token"`
	Student model.Student `json:"student"`
}