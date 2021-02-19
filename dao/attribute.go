package dao

import (
    "database/sql"

    "github.com/firedial/midas-misuzu/entity"
    "github.com/firedial/midas-misuzu/db"
)

type MysqlAttributeRepository struct {
    
}


// eleemnt
func (MysqlAttributeRepository) FindAllElement(attributeName string) (attributes entity.AttributeElements, err error) {
    defer func() { recover() }()

    attributes = []entity.AttributeElement{}

    db := db.Init();
    defer db.Close();

    query := "SELECT " + attributeName + "_element_id AS id, " + attributeName + "_element_name AS name, " + attributeName + "_element_description AS description, " + attributeName + "_category_id AS category_id FROM " + "m_" + attributeName + "_element"

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

// category
func (MysqlAttributeRepository) FindAllCategory(attributeName string) (attributes entity.AttributeCategories, err error) {
    defer func() { recover() }()

    attributes = []entity.AttributeCategory{}

    db := db.Init();
    defer db.Close();

    query := "SELECT " + attributeName + "_category_id AS id, " + attributeName + "_category AS name, " + attributeName + "_category AS description FROM " + "m_" + attributeName + "_category"

    rows, err := db.Query(query)
    defer rows.Close()
    if err != nil {
        return
    }

    for rows.Next() {
        var id int
        var name string
        var description string 

        err := rows.Scan(&id, &name, &description)
        if err != nil {
            return []entity.AttributeCategory{}, err
        }
        attribute := entity.AttributeCategory{
            Id: id,
            Name: name,
            Description: description,
        }
        attributes = append(attributes, attribute)
    }
    return
}

func (MysqlAttributeRepository) SaveCategory(attributeName string, attribute entity.AttributeCategory) (err error) {
    defer func() { recover() }()

    db := db.Init();
    defer db.Close();

    err = saveCategory(attributeName, attribute, db)
    if err != nil {
        return
    }

    return 
}

func saveCategory(attributeName string, attribute entity.AttributeCategory, db *sql.DB) (err error) {

    stmt, err := db.Prepare(`
        INSERT INTO m_` + attributeName + `_category (
            ` + attributeName + `_category_name,
            ` + attributeName + `_category_description
        ) VALUES 
        (?, ?)`)

    if err != nil {
        return 
    }
    defer stmt.Close()

    _, err = stmt.Exec(attribute.Name, attribute.Description)
    if err != nil {
        return
    }

    return
}

// group
func (MysqlAttributeRepository) FindAllGroup(attributeName string) (attributes entity.AttributeGroups, err error) {
    defer func() { recover() }()

    attributes = []entity.AttributeGroup{}

    db := db.Init();
    defer db.Close();

    query := "SELECT " + attributeName + "_group_id AS id, " + attributeName + "_group_name AS name, " + attributeName + "_group_description AS description FROM " + "m_" + attributeName + "_group"

    rows, err := db.Query(query)
    defer rows.Close()
    if err != nil {
        return
    }

    for rows.Next() {
        var id int
        var name string
        var description string 

        err := rows.Scan(&id, &name, &description)
        if err != nil {
            return []entity.AttributeGroup{}, err
        }
        attribute := entity.AttributeGroup{
            Id: id,
            Name: name,
            Description: description,
        }
        attributes = append(attributes, attribute)
    }
    return
}

func (MysqlAttributeRepository) SaveGroup(attributeName string, attribute entity.AttributeGroup) (err error) {
    defer func() { recover() }()

    db := db.Init();
    defer db.Close();

    err = saveGroup(attributeName, attribute, db)
    if err != nil {
        return
    }

    return 
}

func saveGroup(attributeName string, attribute entity.AttributeGroup, db *sql.DB) (err error) {

    stmt, err := db.Prepare(`
        INSERT INTO m_` + attributeName + `_group (
            ` + attributeName + `_group_name,
            ` + attributeName + `_group_description
        ) VALUES 
        (?, ?)`)

    if err != nil {
        return 
    }
    defer stmt.Close()

    _, err = stmt.Exec(attribute.Name, attribute.Description)
    if err != nil {
        return
    }

    return
}

// group_relation
func (MysqlAttributeRepository) FindAllGroupRelation(attributeName string) (attributes entity.AttributeGroupRelations, err error) {
    defer func() { recover() }()

    attributes = []entity.AttributeGroupRelation{}

    db := db.Init();
    defer db.Close();

    query := "SELECT " + attributeName + "_group_id AS group_id, " + attributeName + "_element_id AS element_id FROM " + "m_" + attributeName + "_group_relation"

    rows, err := db.Query(query)
    defer rows.Close()
    if err != nil {
        return
    }

    for rows.Next() {
        var group_id int
        var element_id int

        err := rows.Scan(&group_id, &element_id)
        if err != nil {
            return []entity.AttributeGroupRelation{}, err
        }
        attribute := entity.AttributeGroupRelation{
            GroupId: group_id,
            ElementId: element_id, 
        }
        attributes = append(attributes, attribute)
    }
    return
}

func (MysqlAttributeRepository) SaveGroupRelation(attributeName string, attribute entity.AttributeGroupRelation) (err error) {
    defer func() { recover() }()

    db := db.Init();
    defer db.Close();

    err = saveGroupRelation(attributeName, attribute, db)
    if err != nil {
        return
    }

    return 
}

func saveGroupRelation(attributeName string, attribute entity.AttributeGroupRelation, db *sql.DB) (err error) {

    stmt, err := db.Prepare(`
        INSERT INTO m_` + attributeName + `_group_relation (
            ` + attributeName + `_group_id,
            ` + attributeName + `_element_id
        ) VALUES 
        (?, ?)`)

    if err != nil {
        return 
    }
    defer stmt.Close()

    _, err = stmt.Exec(attribute.GroupId, attribute.ElementId)
    if err != nil {
        return
    }

    return
}


