package repository

import "github.com/firedial/midas-misuzu/entity"

type SumRepository interface {
    Find(map[string][]string) (entity.Sums, error)
}
