package repository

import "github.com/firedial/midas-misuzu/entity"

type AttributeRepository interface {
    FindAll(string) (entity.Attributes, error)
}
