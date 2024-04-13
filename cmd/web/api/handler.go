package api

import "github.com/josemateuss/backend-challenge-frete-rapido/app"

type handler struct {
	application app.Application
}

func New(application app.Application) handler {
	return handler{
		application: application,
	}
}
