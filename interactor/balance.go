package interactor

import (
    "github.com/firedial/midas-go/entity"
    "github.com/firedial/midas-go/repository"
    "github.com/firedial/midas-go/dao"
)

var balanceRepository repository.BalanceRepository = &dao.MysqlBalanceRepository{}

func GetBalance() (entity.Balances, error) {
    return balanceRepository.FindAll()
}

func InsertBalances(balance entity.Balance) string {
    balances := []entity.Balance{balance}

    if !entity.IsSuitableBalances(balances) {
        return "Not Balance!"
    }

    err := balanceRepository.SaveAll(balances);
    if err != nil {
        return "DB error!"
    }

    return "OK"
}
