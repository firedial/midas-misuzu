package controller

import(
    "github.com/firedial/midas-misuzu/entity"
    "github.com/firedial/midas-misuzu/interactor"
)

type returnBalancesJson struct {
    Status string `json:"status"`
    Message string `json:"message"`
    Data entity.Balances `json:"data"`
}

func BalanceGet(queries map[string][]string) returnBalancesJson {
    balances, err := interactor.GetBalance(queries)

    status := "OK"
    message := ""
    if err != nil {
        status = "NG"
        message = err.Error()
    }
    return returnBalancesJson{
        Status: status,
        Message: message,
        Data: balances} 
}

func BalancePost(balance entity.Balance) returnBalancesJson {
    message := interactor.InsertBalances(balance)

    status := "OK"
    if message != "" {
        status = "NG"
    }
    return returnBalancesJson{
        Status: status,
        Message: message,
        Data: []entity.Balance{}}
}
