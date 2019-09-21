package repository

import "github.com/firedial/midas-misuzu/entity"

type SumRepository interface {
    Find(attributeName string, groupByDate string, startDate int, endDate int) (entity.Sums, error)
}
