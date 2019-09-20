package entity

type Attribute struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Description string `json:"description"`
    GroupId int `json:"group_id"`
}

type Attributes []Attribute
