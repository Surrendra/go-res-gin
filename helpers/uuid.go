package helpers

import "github.com/google/uuid"

type uuidHelper struct {
}

func NewUuidHelper() *uuidHelper {
	return &uuidHelper{}
}

type UuidHelper interface {
	GenerateUuid() string
	GenerateTokenForUser() string
}

func (h uuidHelper) GenerateUuid() string {
	//TODO implement me
	return uuid.New().String()
}

func (h uuidHelper) GenerateTokenForUser() string {
	return h.GenerateUuid() + h.GenerateUuid()
}
