package repository

import "github.com/firedial/midas-go/entity"

type SumRepository interface {
    Find(attributeName string, groupByDate string, startDate int, endDate int) (entity.Sums, error)
}
