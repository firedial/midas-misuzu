package interactor

import (
    "strconv"

    "github.com/firedial/midas-go/entity"
    "github.com/firedial/midas-go/repository"
)

func GetSum(attributeName string, groupByDate string, startDate string, endDate string) (entity.Sums, error) {

    intStartDate, _ := strconv.Atoi(startDate)
    intEndDate, _ := strconv.Atoi(endDate)

    return sumRepository.Find(attributeName, groupByDate, intStartDate, intEndDate)
}
