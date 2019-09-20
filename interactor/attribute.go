package interactor

import (
    "github.com/firedial/midas-go/entity"
    "github.com/firedial/midas-go/repository"
)

func GetAttribute(attributeName string) (entity.Attributes, error) {
    return attributeRepository.FindAll(attributeName)
}
