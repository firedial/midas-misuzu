package dao

import (
    "database/sql"

    "github.com/firedial/midas-misuzu/entity"
    "github.com/firedial/midas-misuzu/db"
)

type MysqlBalanceRepository struct {
    
}

func (MysqlBalanceRepository) Find(queries map[string][]string) (balances entity.Balances, err error) {
    defer func() { recover() }()

    balances = []entity.Balance{}

    db := db.Init();
    defer db.Close();

    where, args := getBalanceWhere(queries)

    query := "SELECT balance_id, amount, item, kind_element_id, purpose_element_id, place_element_id, DATE_FORMAT(date, '%Y/%m/%d') as date FROM m_balance"

    rows, err := db.Query(query + where, args...)
    
    defer rows.Close()
    if err != nil {
        return
    }

    for rows.Next() {
        var balance_id int
        var amount int
        var item string
        var kind_id int
        var purpose_id int
        var place_id int
        var date string 

        err := rows.Scan(&balance_id, &amount, &item, &kind_id, &purpose_id, &place_id, &date)
        if err != nil {
            return []entity.Balance{}, err
        }
        balance := entity.Balance{
            BalanceId: balance_id,
            Amount: amount,
            Item: item,
            KindId: kind_id,
            PurposeId: purpose_id,
            PlaceId: place_id,
            Date: date,
        }
        balances = append(balances, balance)
    }
    return
}

func (MysqlBalanceRepository) SaveAll(balances entity.Balances) (err error) {
    db := db.Init();
    defer db.Close();

    tx, err := db.Begin()
    if err != nil {
        return 
    }

    defer func() {
        r := recover()
        if r != nil {
            tx.Rollback()
        }
    }()

    for _, balance := range balances {
        err = save(balance, db)
        if err != nil {
            return
        }
    }

    err = tx.Commit()

    return 
}

func save(balance entity.Balance, db *sql.DB) (err error) {

    stmt, err := db.Prepare(`
        INSERT INTO m_balance (
            amount,
            item,
            kind_element_id,
            purpose_element_id,
            place_element_id,
            date
        ) VALUES 
        (?, ?, ?, ?, ?, ?)`)

    if err != nil {
        return 
    }
    defer stmt.Close()

    _, err = stmt.Exec(balance.Amount, balance.Item, balance.KindId, balance.PurposeId, balance.PlaceId, balance.Date)
    if err != nil {
        return
    }

    return
}

