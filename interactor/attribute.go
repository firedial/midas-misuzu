package interactor

import (
    "github.com/firedial/midas-go/entity"
    "github.com/firedial/midas-go/repository"
    "github.com/firedial/midas-go/dao"
)

var attributeRepository repository.AttributeRepository = &dao.MysqlAttributeRepository{}

func GetAttribute(attributeName string) (entity.Attributes, error) {
    return attributeRepository.FindAll(attributeName)
}
