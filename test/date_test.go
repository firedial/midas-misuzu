package test

import (
    "testing"

    "github.com/firedial/midas-misuzu/util"
)

func TestIsSuitableDate_1(t *testing.T) {
    if util.IsSuitableDate("") {
        t.Error("空文字列は不正です")
    }
    if util.IsSuitableDate("53") {
        t.Error("/が一つもないのは不正です")
    }
    if util.IsSuitableDate("5/3") {
        t.Error("/が1つなのは不正です")
    }
    if util.IsSuitableDate("2019/3/5/5") {
        t.Error("/が3つなのは不正です")
    }
    if util.IsSuitableDate("twa/5/3") {
        t.Error("年が文字列は不正です")
    }
    if util.IsSuitableDate("2019/grag/3") {
        t.Error("月が文字列は不正です")
    }
    if util.IsSuitableDate("2019/5/gra") {
        t.Error("日が文字列は不正です")
    }
    if util.IsSuitableDate("2019.5.3") {
        t.Error("区切り文字が.なのは不正です")
    }
    if !util.IsSuitableDate("2019/5/3") {
        t.Error("正常パターンです")
    }
}

func TestIsSuitableDate_2(t *testing.T) {
    if util.IsSuitableDate("999/5/2") {
        t.Error("年は1000以上の値です")
    }
    if util.IsSuitableDate("10000/5/2") {
        t.Error("年は10000未満の値です")
    }
    if util.IsSuitableDate("4636/0/2") {
        t.Error("月は1以上の値です")
    }
    if util.IsSuitableDate("4636/13/2") {
        t.Error("月は13未満の値です")
    }
    if util.IsSuitableDate("4636/12/0") {
        t.Error("日は1以上の値です")
    }
    if util.IsSuitableDate("4636/12/32") {
        t.Error("日は32未満の値です")
    }
}
