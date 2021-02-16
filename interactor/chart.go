package interactor

import (
    "time"

    "github.com/firedial/midas-misuzu/entity"
)

func GetChartData(queries map[string][]string) (entity.Chart, error) {
    sums, err := sumRepository.Find(queries)

    if err != nil {
        return entity.Chart{}, err
    }

    if len(sums) == 0 {
        return entity.Chart{
            Label: []string{},
            Data: []entity.AttributeData{}}, err
    }

    groupByDate := ""
    if len(queries["groupByDate"]) != 0 {
        groupByDate = queries["groupByDate"][0]
    }
    min, max := getMinMaxDate(sums)
    label := getDateLabel(min, max, groupByDate)

    // attribute の一覧を取得
    attributes := []int{}
    for _, sum := range sums {
        if !isIncludedId(sum.AttributeId, attributes) {
            attributes = append(attributes, sum.AttributeId)
        }
    }

    value := []entity.Sum{}
    data := []entity.AttributeData{}
    for _, attribute := range attributes {
        value = []entity.Sum{}
        for _, sum := range sums {
            if sum.AttributeId == attribute {
                value = append(value, sum)
            }
        }
        attributeData := entity.AttributeData{
            Id: attribute,
            Value: setData(label, value)}
       data = append(data, attributeData) 
    }
    
    return entity.Chart{
        Label: label,
        Data: data}, err
}

func setData(label []string, value []entity.Sum) []int {

    r := []int{}
    for _, d := range label {
        b, i := getIndexDate(d, value)
        if b {
            r = append(r, value[i].AmountSum)
        } else {
            r = append(r, 0)
        }
    }

    return r
}

func getIndexDate(date string, value []entity.Sum) (bool, int) {

    for i, v := range value {
        if v.Date == date {
            return true, i
        }
    }

    return false, 0
}

func isIncludedId(id int, ids []int) bool {

    for _, i := range ids {
        if i == id {
            return true
        }
    }

    return false
}

func getDateLabel(min time.Time, max time.Time, interval string) []string {
    label := []string{}

    y := 1
    m := 0
    d := 0
    switch interval {
    case "day":
        y = 0
        m = 0
        d = 1
    case "week":
        y = 0
        m = 0
        d = 7 
    case "month":
        y = 0
        m = 1 
        d = 0 
    case "year":
        y = 1
        m = 0
        d = 0
    case "fiscal_year":
        y = 1
        m = 0
        d = 0
    default:
        y = 1 
        m = 0
        d = 0
    }

    date := min

    for !max.Before(date) {
        label = append(label, date.Format("2006/01/02"))
        date = date.AddDate(y, m, d)
    }

    return label
}

func getMinMaxDate(sums entity.Sums) (time.Time, time.Time) {
    min, _ := time.Parse("2006/01/02", sums[0].Date)
    max, _ := time.Parse("2006/01/02", sums[0].Date)

    for _, sum := range sums {
        date, _ := time.Parse("2006/01/02", sum.Date)
        if max.Before(date) {
            max = date
        }
        if min.After(date) {
            min = date
        }
    }

    return min, max
}
