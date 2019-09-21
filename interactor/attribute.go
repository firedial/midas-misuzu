package interactor

import (
    "github.com/firedial/midas-misuzu/entity"
)

func GetAttribute(attributeName string) (entity.Attributes, error) {
    return attributeRepository.FindAll(attributeName)
}
