package service

import "context"

type ValidateZipcodeInput struct {
	Zipcode string `json:"zipcode"`
}

type ValidateZipcodeOutput struct {
	Error   bool
	Zipcode string
}

type ValidateZipCode interface {
	Validate(ctx context.Context, input ValidateZipcodeInput) (output *ValidateZipcodeOutput, err error)
}
