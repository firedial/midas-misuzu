package controller

import(
    "github.com/firedial/midas-misuzu/entity"
    "github.com/firedial/midas-misuzu/interactor"
)

func ChartGet(queries map[string][]string) entity.Chart {
    /*
    r := entity.Chart{
        Label: [] string{"a", "b"},
        Data: [] entity.AttributeData{
            entity.AttributeData{
                Id: 1,
                Value: [] int{1000, 2000}},
            entity.AttributeData{
                Id: 2,
                Value: [] int{1500, 2500}}}}
    */

    data, _ := interactor.GetChartData(queries)
    return data
}
