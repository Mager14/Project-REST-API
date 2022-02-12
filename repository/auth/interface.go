package auth

import "Project-REST-API/entities"

type Auth interface {
	Login(name, hp string) (entities.User, error)
}
