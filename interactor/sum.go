package interactor

import (
    "github.com/firedial/midas-misuzu/entity"
)

func GetSum(queries map[string][]string) (entity.Sums, error) {
    return sumRepository.Find(queries)
}
