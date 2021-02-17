package controller

import(
    "github.com/firedial/midas-misuzu/interactor"
    "github.com/firedial/midas-misuzu/entity"
)

type returnAttributesJson struct {
    Status string `json:"status"`
    Message string `json:"message"`
    Data entity.Attributes `json:"data"`
}


func AttributeElementsGet(attribute_name string) returnAttributesJson {
    attributes, err := interactor.GetAttributeElements(attribute_name)
    
    status := "OK"
    message := ""
    if err != nil {
        status = "NG"
        message = err.Error()
    }
    return returnAttributesJson{
        Status: status,
        Message: message,
        Data: attributes} 
}

func AttributeElementPost(attributeName string, attribute entity.Attribute) returnAttributesJson {
    message := interactor.InsertAttributeElement(attributeName, attribute)

    status := "OK"
    if message != "" {
        status = "NG"
    }
    return returnAttributesJson{
        Status: status,
        Message: message,
        Data: []entity.Attribute{}}
}
