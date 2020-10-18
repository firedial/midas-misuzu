package entity

type Attribute struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Description string `json:"description"`
    CategoryId int `json:"category_id"`
}

type Attributes []Attribute
