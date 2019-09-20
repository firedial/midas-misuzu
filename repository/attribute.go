package repository

import "github.com/firedial/midas-go/entity"

type AttributeRepository interface {
    FindAll(string) (entity.Attributes, error)
}
