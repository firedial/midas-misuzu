package controller

import(
    "github.com/firedial/midas-misuzu/interactor"
    "github.com/firedial/midas-misuzu/entity"
)

func AttributeGet(attribute_name string) entity.Attributes {
    attributes, _ := interactor.GetAttribute(attribute_name)
    return attributes
}
