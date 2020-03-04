package repository

import "github.com/firedial/midas-misuzu/entity"

type BalanceRepository interface {
    SaveAll(entity.Balances) error
    FindAll() (entity.Balances, error)
    Find([]string) (entity.Balances, error)
}
