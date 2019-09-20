package repository

import "github.com/firedial/midas-go/entity"

type BalanceRepository interface {
    SaveAll(entity.Balances) error
    FindAll() (entity.Balances, error)
}
