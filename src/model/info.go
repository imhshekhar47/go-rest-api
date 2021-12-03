package model

type Info struct {
	Name string `json:"name" binding:"required"`
	Version string `json:"version" binding:"required"`
}
