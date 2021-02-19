package repository

import "github.com/firedial/midas-misuzu/entity"

type AttributeRepository interface {
    FindAllElement(string) (entity.AttributeElements, error)
    SaveElement(string, entity.AttributeElement) error
}
