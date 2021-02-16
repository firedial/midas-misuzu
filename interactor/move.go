package interactor

import (
    "github.com/firedial/midas-misuzu/entity"
    "github.com/firedial/midas-misuzu/config"
)

type Move struct {
    Attribute string `json:"attribute"`
    Amount int `json:"amount"`
    BeforeId int `json:"before_id"`
    AfterId int `json:"after_id"`
    Date string `json:"date"`
}

func InsertMove(move Move) string {

    balances := getMoveBalance(move)

    if !entity.IsSuitableBalances(balances) {
        return "\"Not Move!\""
    }

    err := balanceRepository.SaveAll(balances);
    if err != nil {
        return err.Error()
    }

    return ""
}

func getMoveBalance(move Move) entity.Balances {
    var beforeBalance entity.Balance
    var afterBalance entity.Balance

    beforeBalance.Amount = -1 * move.Amount
    afterBalance.Amount = move.Amount
    beforeBalance.KindId = config.KIND_MOVE_ID
    afterBalance.KindId = config.KIND_MOVE_ID
    beforeBalance.Date = move.Date
    afterBalance.Date = move.Date

    if move.Attribute == "purpose" {
        beforeBalance.Item = "予算移動元"
        afterBalance.Item = "予算移動先"
        beforeBalance.PlaceId = config.PLACE_MOVE_ID 
        afterBalance.PlaceId = config.PLACE_MOVE_ID

        beforeBalance.PurposeId = move.BeforeId 
        afterBalance.PurposeId = move.AfterId
    } else if move.Attribute == "place" {
        beforeBalance.Item = "場所移動元"
        afterBalance.Item = "場所移動先"
        beforeBalance.PurposeId = config.PURPOSE_MOVE_ID
        afterBalance.PurposeId = config.PURPOSE_MOVE_ID

        beforeBalance.PlaceId = move.BeforeId 
        afterBalance.PlaceId = move.AfterId
    } else {
        panic("non")
    }

    return entity.Balances{beforeBalance, afterBalance}
}
