package controller

import(
    "github.com/firedial/midas-go/interactor"
    "github.com/firedial/midas-go/entity"
)

func AttributeGet(attribute_name string) entity.Attributes {
    attributes, _ := interactor.GetAttribute(attribute_name)
    return attributes
}
