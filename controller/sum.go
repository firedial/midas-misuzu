package controller

import(
    "github.com/firedial/midas-misuzu/interactor"
    "github.com/firedial/midas-misuzu/entity"
)

type returnSumsJson struct {
    Status string `json:"status"`
    Message string `json:"message"`
    Data entity.Sums `json:"data"`
}

func SumGet(queries map[string][]string) returnSumsJson {
    sums, err := interactor.GetSum(queries)

    message := ""
    status := "OK"
    if err != nil {
        status = "NG"
        message = err.Error()
    }
    return returnSumsJson{
        Status: status,
        Message: message,
        Data: sums}
}

