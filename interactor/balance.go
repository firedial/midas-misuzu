package interactor

import (
    "github.com/firedial/midas-misuzu/entity"
    "github.com/firedial/midas-misuzu/config"
)

func GetBalance(queries map[string][]string) (entity.Balances, error) {
    return balanceRepository.Find(queries)
}

func InsertBalances(balance entity.Balance) string {

    if balance.KindId == config.KIND_MOVE_ID {
        return "\"Can not use kind move id!\""
    }

    if balance.PurposeId == config.PURPOSE_MOVE_ID {
        return "\"Can not use purpose move id!\""
    }

    if balance.PlaceId == config.PLACE_MOVE_ID {
        return "\"Can not use place move id!\""
    }

    balances := []entity.Balance{balance}

    if !entity.IsSuitableBalances(balances) {
        return "\"Not Balance!\""
    }

    err := balanceRepository.SaveAll(balances);
    if err != nil {
        return err.Error()
    }

    return ""
}
