package repository

import "github.com/firedial/midas-misuzu/entity"

type AttributeRepository interface {
    FindAllElement(string) (entity.AttributeElements, error)
    SaveElement(string, entity.AttributeElement) error

    FindAllCategory(string) (entity.AttributeCategories, error)
    SaveCategory(string, entity.AttributeCategory) error

    FindAllGroup(string) (entity.AttributeGroups, error)
    SaveGroup(string, entity.AttributeGroup) error

    FindAllGroupRelation(string) (entity.AttributeGroupRelations, error)
    SaveGroupRelation(string, entity.AttributeGroupRelation) error
}
