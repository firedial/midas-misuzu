package controller

import(
    "github.com/firedial/midas-misuzu/interactor"
    "github.com/firedial/midas-misuzu/entity"
)

type returnAttributeElementsJson struct {
    Status string `json:"status"`
    Message string `json:"message"`
    Data entity.AttributeElements `json:"data"`
}

type returnAttributeCategoriesJson struct {
    Status string `json:"status"`
    Message string `json:"message"`
    Data entity.AttributeCategories `json:"data"`
}

type returnAttributeGroupsJson struct {
    Status string `json:"status"`
    Message string `json:"message"`
    Data entity.AttributeGroups `json:"data"`
}

type returnAttributeGroupRelationsJson struct {
    Status string `json:"status"`
    Message string `json:"message"`
    Data entity.AttributeGroupRelations `json:"data"`
}


func AttributeElementsGet(attribute_name string) returnAttributeElementsJson {
    attributes, err := interactor.GetAttributeElements(attribute_name)
    
    status := "OK"
    message := ""
    if err != nil {
        status = "NG"
        message = err.Error()
    }
    return returnAttributeElementsJson{
        Status: status,
        Message: message,
        Data: attributes} 
}

func AttributeElementPost(attributeName string, attribute entity.AttributeElement) returnAttributeElementsJson {
    message := interactor.InsertAttributeElement(attributeName, attribute)

    status := "OK"
    if message != "" {
        status = "NG"
    }
    return returnAttributeElementsJson{
        Status: status,
        Message: message,
        Data: []entity.AttributeElement{}}
}

func AttributeCategoriesGet(attribute_name string) returnAttributeCategoriesJson {
    attributes, err := interactor.GetAttributeCategories(attribute_name)
    
    status := "OK"
    message := ""
    if err != nil {
        status = "NG"
        message = err.Error()
    }
    return returnAttributeCategoriesJson{
        Status: status,
        Message: message,
        Data: attributes} 
}

func AttributeCategoryPost(attributeName string, attribute entity.AttributeCategory) returnAttributeCategoriesJson {
    message := interactor.InsertAttributeCategory(attributeName, attribute)

    status := "OK"
    if message != "" {
        status = "NG"
    }
    return returnAttributeCategoriesJson{
        Status: status,
        Message: message,
        Data: []entity.AttributeCategory{}}
}

func AttributeGroupsGet(attribute_name string) returnAttributeGroupsJson {
    attributes, err := interactor.GetAttributeGroups(attribute_name)
    
    status := "OK"
    message := ""
    if err != nil {
        status = "NG"
        message = err.Error()
    }
    return returnAttributeGroupsJson{
        Status: status,
        Message: message,
        Data: attributes} 
}

func AttributeGroupPost(attributeName string, attribute entity.AttributeGroup) returnAttributeGroupsJson {
    message := interactor.InsertAttributeGroup(attributeName, attribute)

    status := "OK"
    if message != "" {
        status = "NG"
    }
    return returnAttributeGroupsJson{
        Status: status,
        Message: message,
        Data: []entity.AttributeGroup{}}
}

func AttributeGroupRelationsGet(attribute_name string) returnAttributeGroupRelationsJson {
    attributes, err := interactor.GetAttributeGroupRelations(attribute_name)
    
    status := "OK"
    message := ""
    if err != nil {
        status = "NG"
        message = err.Error()
    }
    return returnAttributeGroupRelationsJson{
        Status: status,
        Message: message,
        Data: attributes} 
}

func AttributeGroupRelationPost(attributeName string, attribute entity.AttributeGroupRelation) returnAttributeGroupRelationsJson {
    message := interactor.InsertAttributeGroupRelation(attributeName, attribute)

    status := "OK"
    if message != "" {
        status = "NG"
    }
    return returnAttributeGroupRelationsJson{
        Status: status,
        Message: message,
        Data: []entity.AttributeGroupRelation{}}
}

