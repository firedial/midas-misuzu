package controller

import(
    "github.com/firedial/midas-misuzu/interactor"
)

type returnMoveJson struct {
    Status string `json:"status"`
    Message string `json:"message"`
}

func MovePost(move interactor.Move) returnMoveJson {
    message := interactor.InsertMove(move)

    status := "OK"
    if message != "" {
        status = "NG"
    }
    return returnMoveJson{
        Status: status,
        Message: message}
}
