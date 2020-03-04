package controller

import(
    "github.com/firedial/midas-misuzu/entity"
    "github.com/firedial/midas-misuzu/interactor"
)

func BalanceGet(queries map[string][]string) entity.Balances {
    balances, _ := interactor.GetBalance(queries)
    return balances
}

func BalancePost(balance entity.Balance) string {
    return interactor.InsertBalances(balance)
}
