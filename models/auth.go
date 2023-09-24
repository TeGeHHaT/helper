package models

type UserAuth struct {
	Name     string `json:"user"`
	Password string `json:"password"`
}
