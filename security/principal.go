package security

import "github.com/say-bonjour/be-say-bonjour/model"

type Principal struct {
	Id    uint64
	Name  string
	Email string
	Role  model.Role
}
