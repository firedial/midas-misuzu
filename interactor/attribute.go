package interactor

import (
    "github.com/firedial/midas-misuzu/entity"
)

func GetAttributeElements(attributeName string) (entity.Attributes, error) {
    return attributeRepository.FindAllElement(attributeName)
}

func InsertAttributeElement(attributeName string, attribute entity.Attribute) string {
    err := attributeRepository.SaveElement(attributeName, attribute)
    
    if err != nil {
        return err.Error()
    }

    return ""
}
