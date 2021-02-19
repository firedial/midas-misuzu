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
