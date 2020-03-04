package controller

import(
    "github.com/firedial/midas-misuzu/interactor"
    "github.com/firedial/midas-misuzu/entity"
)

func SumGet(queries map[string][]string) entity.Sums {
    sums, _ := interactor.GetSum(queries)
    return sums
}

func getQuery(queries map[string][]string, name string) string {
    vals, ok := queries[name]
    if !ok {
        return ""
    }

    return vals[0]
}
