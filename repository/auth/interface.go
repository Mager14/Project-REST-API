package auth

import "Project-REST-API/entities"

type Auth interface {
	Login(email, password string) (entities.User, error)
}
