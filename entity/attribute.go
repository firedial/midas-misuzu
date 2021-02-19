package entity

type AttributeElement struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Description string `json:"description"`
    CategoryId int `json:"category_id"`
}

type AttributeElements []AttributeElement

type AttributeCategory struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Description string `json:"description"`
}

type AttributeCategories []AttributeCategory

type AttributeGroup struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Description string `json:"description"`
}

type AttributeGroups []AttributeGroup

type AttributeGroupRelation struct {
    GroupId int `json:"group_id"`
    ElementId int `json:"element_id"`
}

type AttributeGroupRelations []AttributeGroupRelation
