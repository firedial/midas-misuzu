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


func AttributeGet(attribute_name string) returnAttributesJson {
    attributes, err := interactor.GetAttribute(attribute_name)
    
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
