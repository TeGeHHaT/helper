package models

type Direction struct {
	Id         int    `json:"id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	Color      string `json:"color"`
	IsDisabled bool   `json:"is_disabled"`
}

type DirectionGetParams struct {
	Id int `json:"id"`
}

type DirectionInsUpdParams struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type DirectionDelParams struct {
	Ids string `json:"ids"`
}
