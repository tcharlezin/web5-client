package models

type User struct {
	Uuid      string `json:"uuid"`
	PublicKey string `json:"public_key"`
}
