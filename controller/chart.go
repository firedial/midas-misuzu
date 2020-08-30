package controller

import(
    "github.com/firedial/midas-misuzu/entity"
)

func ChartGet(queries map[string][]string) entity.Chart {
    r := entity.Chart{
        Label: [] string{"a", "b"},
        Data: [] entity.AttributeData{
            entity.AttributeData{
                Name: "lunch",
                Value: [] int{1000, 2000}},
            entity.AttributeData{
                Name: "dinner",
                Value: [] int{1500, 2500}}}}
    
    return r
}
