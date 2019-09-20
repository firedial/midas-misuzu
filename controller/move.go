package controller

import(
    "github.com/firedial/midas-go/interactor"
)

func MovePost(move interactor.Move) string {
    return interactor.InsertMove(move)
}
