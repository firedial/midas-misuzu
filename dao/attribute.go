package dao

import (
    "github.com/firedial/midas-misuzu/entity"
    "github.com/firedial/midas-misuzu/db"
)

type MysqlAttributeRepository struct {
    
}

func (MysqlAttributeRepository) FindAll(attributeName string) (attributes entity.Attributes, err error) {
    db := db.Init();
    defer db.Close();

    query := "SELECT " + attributeName + "_element_id AS id, " + attributeName + "_element_name, " + attributeName + "_element_description, " + attributeName + "_category_id FROM " + "m_" + attributeName + "_element"

    rows, err := db.Query(query)
    defer rows.Close()
    if err != nil {
        return
    }

    for rows.Next() {
        var id int
        var name string
        var description string 
        var category_id int

        err := rows.Scan(&id, &name, &description, &category_id)
        if err != nil {
            return []entity.Attribute{}, err
        }
        attribute := entity.Attribute{
            Id: id,
            Name: name,
            Description: description,
            CategoryId: category_id,
        }
        attributes = append(attributes, attribute)
    }
    return
}
