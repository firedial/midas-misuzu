package entity

type AttributeData struct {
    Id int `json:"id"` 
    Value []int `json:"value"`
}

type Chart struct {
    Label []string `json:"label"`
    Data []AttributeData `json:"data"`
}
