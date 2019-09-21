package entity

import (
    "github.com/firedial/midas-misuzu/util"
)

type Balance struct {
    BalanceId int `json:"balance_id"`
    Amount int `json:"amount"`
    Item string `json:"item"`
    KindId int `json:"kind_id"`
    PurposeId int `json:"purpose_id"`
    PlaceId int `json:"place_id"`
    Date string `json:"date"`
}

type Balances []Balance

func IsSuitableBalances(balances Balances) bool {
    for _, balance := range balances {
        if !isSuitableBalance(balance) {
            return false
        }
    }

    return true
}

func isSuitableBalance(balance Balance) bool {
    if balance.Amount == 0 {
        return false
    }
    if !util.IsSuitableString(balance.Item) {
        return false
    }
    if balance.KindId <= 0 {
        return false
    }
    if balance.PurposeId <= 0 {
        return false
    }
    if balance.PlaceId <= 0 {
        return false
    }
    if !util.IsSuitableDate(balance.Date) {
        return false
    }

    return true
}
