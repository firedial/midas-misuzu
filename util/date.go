package util

import (
    "strings"
    "strconv"
)

func IsSuitableDate(str string) bool {
    splitStr := strings.Split(str, "/")

    // "/" で 3 つに分割できない場合は不適
    if len(splitStr) != 3 {
        return false
    }

    year, erry := strconv.Atoi(splitStr[0])
    month, errm := strconv.Atoi(splitStr[1])
    day, errd := strconv.Atoi(splitStr[2])

    if erry != nil || errm != nil || errd != nil {
        return false
    }

    // 年が 1000 から 9999 の間であるかどうか
    if !(1000 <= year && year <= 9999) {
        return false
    }

    // 月が 1 から 12 の間であるかどうか
    if !(1 <= month && month <= 12) {
        return false
    }

    // 日が 1 から 31 の間であるかどうか
    // 31日がない月もあるがそこまでは考えない
    if !(1 <= day && day <= 31) {
        return false
    }

    return true
}
