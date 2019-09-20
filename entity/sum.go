package entity

type Sum struct {
    Id int `json:"id"`
    Date string `json:"date"`
    AmountSum int `json:"amount_sum"`
}

type Sums []Sum
