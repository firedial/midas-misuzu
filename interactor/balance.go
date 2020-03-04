package interactor

import (
    "github.com/firedial/midas-misuzu/entity"
)

func GetBalance(queries map[string][]string) (entity.Balances, error) {
    return balanceRepository.Find(queries)
}

func InsertBalances(balance entity.Balance) string {
    balances := []entity.Balance{balance}

    if !entity.IsSuitableBalances(balances) {
        return "\"Not Balance!\""
    }

    err := balanceRepository.SaveAll(balances);
    if err != nil {
        return "\"DB error!\""
    }

    return "\"OK\""
}
