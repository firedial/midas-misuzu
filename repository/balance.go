package repository

import "github.com/firedial/midas-misuzu/entity"

type BalanceRepository interface {
    SaveAll(entity.Balances) error
    Find(map[string][]string) (entity.Balances, error)
}
