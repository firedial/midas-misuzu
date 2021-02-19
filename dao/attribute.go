package dao

import (
    "database/sql"

    "github.com/firedial/midas-misuzu/entity"
    "github.com/firedial/midas-misuzu/db"
)

type MysqlAttributeRepository struct {
    
}

func (MysqlAttributeRepository) FindAllElement(attributeName string) (attributes entity.AttributeElements, err error) {
    defer func() { recover() }()

    attributes = []entity.AttributeElement{}

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
            return []entity.AttributeElement{}, err
        }
        attribute := entity.AttributeElement{
            Id: id,
            Name: name,
            Description: description,
            CategoryId: category_id,
        }
        attributes = append(attributes, attribute)
    }
    return
}

func (MysqlAttributeRepository) SaveElement(attributeName string, attribute entity.AttributeElement) (err error) {
    defer func() { recover() }()

    db := db.Init();
    defer db.Close();

    err = saveElement(attributeName, attribute, db)
    if err != nil {
        return
    }

    return 
}

func saveElement(attributeName string, attribute entity.AttributeElement, db *sql.DB) (err error) {

    stmt, err := db.Prepare(`
        INSERT INTO m_` + attributeName + `_element (
            ` + attributeName + `_element_name,
            ` + attributeName + `_element_description,
            ` + attributeName + `_category_id
        ) VALUES 
        (?, ?, ?)`)

    if err != nil {
        return 
    }
    defer stmt.Close()

    _, err = stmt.Exec(attribute.Name, attribute.Description, attribute.CategoryId)
    if err != nil {
        return
    }

    return
}


