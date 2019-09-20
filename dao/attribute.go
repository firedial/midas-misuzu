package dao

import (
    "github.com/firedial/midas-go/entity"
    "github.com/firedial/midas-go/db"
)

type MysqlAttributeRepository struct {
    
}

func (MysqlAttributeRepository) FindAll(attributeName string) (attributes entity.Attributes, err error) {
    db := db.Init();
    defer db.Close();

    rows, err := db.Query("SELECT id, name, description, group_id FROM " + attributeName)
    defer rows.Close()
    if err != nil {
        return
    }

    for rows.Next() {
        var id int
        var name string
        var description string 
        var group_id int

        err := rows.Scan(&id, &name, &description, &group_id)
        if err != nil {
            return []entity.Attribute{}, err
        }
        attribute := entity.Attribute{
            Id: id,
            Name: name,
            Description: description,
            GroupId: group_id,
        }
        attributes = append(attributes, attribute)
    }
    return
}
