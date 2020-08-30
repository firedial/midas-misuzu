package entity

type AttributeData struct {
    Name string `json:"name"` 
    Value []int `json:"value"`
}

type Chart struct {
    Label []string `json:"label"`
    Data []AttributeData `json:"data"`
}
