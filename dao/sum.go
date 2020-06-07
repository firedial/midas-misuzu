package dao

import (
    "github.com/firedial/midas-misuzu/entity"
    "github.com/firedial/midas-misuzu/db"
    "strconv"
)

type MysqlSumRepository struct {
    
}

func (MysqlSumRepository) Find(queries map[string][]string) (sums entity.Sums, err error) {
    db := db.Init();
    defer db.Close();

    where, args := getBalanceWhere(queries)

    var groupByDateQuery string
    var dateColumn string
    groupByDate := ""
    if len(queries["groupByDate"]) != 0 {
        groupByDate = queries["groupByDate"][0]
    }
    switch groupByDate {
    case "day":
        groupByDateQuery = "DATE_FORMAT(date, '%Y/%m/%d')"
        dateColumn = "DATE_FORMAT(date, '%Y/%m/%d') as date"
    case "week":
        groupByDateQuery = "DATE_FORMAT(SUBDATE(date, WEEKDAY(date)), '%Y/%m/%d')"
        dateColumn = "DATE_FORMAT(SUBDATE(date, WEEKDAY(date)), '%Y/%m/%d') as date"
    case "month":
        groupByDateQuery = "DATE_FORMAT(date, '%Y/%m/01')"
        dateColumn = "DATE_FORMAT(date, '%Y/%m/01') as date"
    case "year":
        groupByDateQuery = "DATE_FORMAT(date, '%Y/01/01')"
        dateColumn = "DATE_FORMAT(date, '%Y/01/01') as date"
    case "fiscal_year":
        groupByDateQuery = "DATE_FORMAT(date - INTERVAL 3 MONTH, '%Y/04/01')"
        dateColumn = "DATE_FORMAT(date - INTERVAL 3 MONTH, '%Y/04/01') as date"
    default:
        groupByDateQuery = ""
        dateColumn = "\"1000/1/1\" as date"
    }

    var groupByAttributeQuery string
    var attributeColumn string
    attributeName := ""
    if len(queries["attributeName"]) != 0 {
        attributeName = queries["attributeName"][0]
    }
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

    query := "SELECT " + attributeColumn + ", " + dateColumn + ", " + "sum(amount) as sum FROM balance"

    rows, err := db.Query(query + " " + where + " " + groupBy, args...)
    defer rows.Close()
    if err != nil {
        return []entity.Sum{}, err
    }

    for rows.Next() {
        var attributeId int
        var date string 
        var amountSum int

        err := rows.Scan(&attributeId, &date, &amountSum)
        if err != nil {
            return []entity.Sum{}, err
        }
        sum := entity.Sum{
            Id: (strconv.Itoa(attributeId) + "-" + date),
            AttributeId: attributeId,
            Date: date,
            AmountSum: amountSum,
        }
        sums = append(sums, sum)
    }
    return
}
