package controller

import(
    "github.com/firedial/midas-misuzu/interactor"
)

func MovePost(move interactor.Move) string {
    return interactor.InsertMove(move)
}
