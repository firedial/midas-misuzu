package util

import (
    "strings"
)

// 文字列の中に['":. ]の5種類の半角記号が入っていないことを見る
// @return bool true: それらが入っていない / false: それらが入っている
func IsSuitableString(str string) bool {
    checkChars := []string{"'", "\"", ":", ".", " "}
    
    if str == "" {
        return false
    }

    for _, c := range checkChars {
        if strings.Contains(str, c) {
            return false
        }
    }

    return true
}
