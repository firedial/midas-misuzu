package dao

import (
    "strconv"
    
    "github.com/firedial/midas-go/entity"
    "github.com/firedial/midas-go/db"
)

type MysqlSumRepository struct {
    
}

func (MysqlSumRepository) Find(attributeName string, groupByDate string, startDate int, endDate int) (sums entity.Sums, err error) {
    db := db.Init();
    defer db.Close();

    var groupByDateQuery string
    var dateColumn string
    switch groupByDate {
    case "day":
        groupByDateQuery = "DATE_FORMAT(date, '%Y%m%d')"
        dateColumn = "DATE_FORMAT(date, '%Y%m%d') as date"
    case "month":
        groupByDateQuery = "DATE_FORMAT(date, '%Y%m')"
        dateColumn = "DATE_FORMAT(date, '%Y%m') as date"
    case "year":
        groupByDateQuery = "DATE_FORMAT(date, '%Y')"
        dateColumn = "DATE_FORMAT(date, '%Y') as date"
    case "fiscal_year":
        groupByDateQuery = "DATE_FORMAT(date - INTERVAL 3 MONTH, '%Y')"
        dateColumn = "DATE_FORMAT(date - INTERVAL 3 MONTH, '%Y') as date"
    default:
        groupByDateQuery = ""
        dateColumn = "\"1000/1/1\" as date"
    }

    var groupByAttributeQuery string
    var attributeColumn string
    switch attributeName {
    case "place":
        groupByAttributeQuery = "place_id"
        attributeColumn = "place_id as id"
    case "purpose":
        groupByAttributeQuery = "purpose_id"
        attributeColumn = "purpose_id as id"
    case "kind":
        groupByAttributeQuery = "kind_id"
        attributeColumn = "kind_id as id"
    default:
        groupByAttributeQuery = ""
        attributeColumn = "0 as id"
    }

    var groupBy string
    if groupByDateQuery != "" {
        if groupByAttributeQuery != "" {
            groupBy = "GROUP BY " + groupByDateQuery + ", " + groupByAttributeQuery
        } else {
            groupBy = "GROUP BY " + groupByDateQuery 
        }
    } else {
        if groupByAttributeQuery != "" {
            groupBy = "GROUP BY " + groupByAttributeQuery
        } else {
            groupBy = "" 
        }
    }

    where := "WHERE 1 = 1"
    if startDate != 0 {
        where += " AND date >= " + strconv.Itoa(startDate)
    }
    if endDate != 0 {
        where += " AND date <= " + strconv.Itoa(endDate)
    }

    query := "SELECT " + attributeColumn + ", " + dateColumn + ", " + "sum(amount) as sum FROM balance"

    rows, err := db.Query(query + " " + where + " " + groupBy)
    defer rows.Close()
    if err != nil {
        return []entity.Sum{}, err
    }

    for rows.Next() {
        var id int
        var date string 
        var amountSum int

        err := rows.Scan(&id, &date, &amountSum)
        if err != nil {
            return []entity.Sum{}, err
        }
        sum := entity.Sum{
            Id: id,
            Date: date,
            AmountSum: amountSum,
        }
        sums = append(sums, sum)
    }
    return
}
