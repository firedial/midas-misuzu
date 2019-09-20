package controller

import(
    "github.com/firedial/midas-go/interactor"
    "github.com/firedial/midas-go/entity"
)

func SumGet(queries map[string][]string) entity.Sums {
    sums, _ := interactor.GetSum(
        getQuery(queries, "attributeName"),
        getQuery(queries, "groupByDate"),
        getQuery(queries, "startDate"),
        getQuery(queries, "endDate"))
    return sums
}

func getQuery(queries map[string][]string, name string) string {
    vals, ok := queries[name]
    if !ok {
        return ""
    }

    return vals[0]
}
