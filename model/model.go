package model

import (
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Author struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" query:"id,omitempty" url:"_id,omitempty" reqHeader:"id"`
	Profile   string             `json:"profile,omitempty" bson:"profile,omitempty"`
	Nama      string             `json:"nama,omitempty" bson:"nama,omitempty" validate:"required,min=2"`
	NIK       string             `json:"nik,omitempty" bson:"nik,omitempty" validate:"required,number,min=16"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty" validate:"required,email,min=6"`
	Phone     string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Pekerjaan string             `json:"pekerjaan,omitempty" bson:"pekerjaan,omitempty" validate:"required,number"`
	Alamat    string             `json:"alamat,omitempty" bson:"alamat,omitempty" validate:"required,number"`
	Bio       string             `json:"bio,omitempty" bson:"bio,omitempty"`
	Photo     string             `json:"photo,omitempty" bson:"photo,omitempty"`
}

type Login struct {
	Login string `json:"login,omitempty" bson:"login,omitempty" query:"login" url:"login,omitempty" reqHeader:"login"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateAuthorStruct(user Author) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
