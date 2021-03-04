package controller

import(
    "github.com/firedial/midas-misuzu/auth"
)

type returnAuthJson struct {
    Status string `json:"status"`
    Message string `json:"message"`
    Data string `json:"data"`
}

func AuthPost(salt string, pass string) returnAuthJson {
    token := auth.GetAuthToken(salt, pass)

    return returnAuthJson{
        Status: "OK",
        Message: "",
        Data: token}
}

func AuthCheck(isSame bool) returnAuthJson {
    status := "OK"
    message := ""

    if !isSame {
        status = "NG"
        message = "Invalid auth token!"
    }

    return returnAuthJson{
        Status: status,
        Message: message,
        Data: ""}
}
