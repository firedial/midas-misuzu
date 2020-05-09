package entity

type Sum struct {
    Id string `json:"id"`
    AttributeId int `json:"attribute_id"`
    Date string `json:"date"`
    AmountSum int `json:"amount_sum"`
}

type Sums []Sum
