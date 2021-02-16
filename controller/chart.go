package controller

import(
    "github.com/firedial/midas-misuzu/entity"
    "github.com/firedial/midas-misuzu/interactor"
)

type returnChartJson struct {
    Status string `json:"status"`
    Message string `json:"message"`
    Data entity.Chart `json:"data"`
}

func ChartGet(queries map[string][]string) returnChartJson {
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

    data, err := interactor.GetChartData(queries)

    message := ""
    status := "OK"
    if err != nil {
        status = "NG"
        message = err.Error()
    }
    return returnChartJson{
        Status: status,
        Message: message,
        Data: data}
}
