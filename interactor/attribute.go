package interactor

import (
    "github.com/firedial/midas-misuzu/entity"
)

func GetAttributeElements(attributeName string) (entity.AttributeElements, error) {
    return attributeRepository.FindAllElement(attributeName)
}

func InsertAttributeElement(attributeName string, attribute entity.AttributeElement) string {
    err := attributeRepository.SaveElement(attributeName, attribute)
    
    if err != nil {
        return err.Error()
    }

    return ""
}

func GetAttributeCategories(attributeName string) (entity.AttributeCategories, error) {
    return attributeRepository.FindAllCategory(attributeName)
}

func InsertAttributeCategory(attributeName string, attribute entity.AttributeCategory) string {
    err := attributeRepository.SaveCategory(attributeName, attribute)
    
    if err != nil {
        return err.Error()
    }

    return ""
}

func GetAttributeGroups(attributeName string) (entity.AttributeGroups, error) {
    return attributeRepository.FindAllGroup(attributeName)
}

func InsertAttributeGroup(attributeName string, attribute entity.AttributeGroup) string {
    err := attributeRepository.SaveGroup(attributeName, attribute)
    
    if err != nil {
        return err.Error()
    }

    return ""
}

func GetAttributeGroupRelations(attributeName string) (entity.AttributeGroupRelations, error) {
    return attributeRepository.FindAllGroupRelation(attributeName)
}

func InsertAttributeGroupRelation(attributeName string, attribute entity.AttributeGroupRelation) string {
    err := attributeRepository.SaveGroupRelation(attributeName, attribute)
    
    if err != nil {
        return err.Error()
    }

    return ""
}
