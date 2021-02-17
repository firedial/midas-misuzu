package repository

import "github.com/firedial/midas-misuzu/entity"

type AttributeRepository interface {
    FindAllElement(string) (entity.Attributes, error)
    SaveElement(string, entity.Attribute) error
}
