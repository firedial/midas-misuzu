package dao


func getBalanceWhere(queries map[string][]string) (string, []interface{}) {
    args := []interface{}{}

    whereKind := ""
    if len(queries["kind"]) != 0 {
        whereKind = " and kind_id in " + getPlaceholder(len(queries["kind"]))
        args = pushArgs(args, queries["kind"])
    }
    wherePurpose := ""
    if len(queries["purpose"]) != 0 {
        wherePurpose = " and purpose_id in " + getPlaceholder(len(queries["purpose"]))
        args = pushArgs(args, queries["purpose"])
    }
    wherePlace := ""
    if len(queries["place"]) != 0 {
        wherePlace = " and place_id in " + getPlaceholder(len(queries["place"]))
        args = pushArgs(args, queries["place"])
    }

    whereStart := ""
    if len(queries["start"]) != 0 {
        whereStart = " and date >= ?"
        args = append(args, queries["start"][0])
    }
    whereEnd := ""
    if len(queries["end"]) != 0 {
        whereEnd = " and date <= ?"
        args = append(args, queries["end"][0])
    }

    whereMove := ""
    if len(queries["move"]) != 0 {
        whereMove = " and kind_id <> 0 and purpose_id <> 0 and place_id <> 0"
    }

    where := " where 1 = 1" + whereKind + wherePurpose + wherePlace + whereStart + whereEnd + whereMove

    return where, args

}

func getPlaceholder(n int) string {
    p := "(?"
    for i := 1; i < n; i++ {
        p += ", ?"
    }
    return p + ")"
}

func pushArgs(args []interface{}, vals []string) []interface{} {
    for _, val := range vals {
        args = append(args, val)
    }
    return args
}

