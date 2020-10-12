package cerealMessages

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (x *Decimal) ToBson() primitive.Decimal128 {
	return primitive.NewDecimal128(x.High, x.Low)
}

func DecimalFromBson(value primitive.Decimal128) *Decimal {
	high, low := value.GetBytes()
	return &Decimal{
		High: high,
		Low:  low,
	}
}
