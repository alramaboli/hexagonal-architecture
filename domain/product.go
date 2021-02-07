package product

//Product ...
type Product struct {
	Code  string  `json:"code" bson:"code"`
	Name  string  `json:"name" bson:"name"`
	Price float32 `json:"price" bson:"price"`
}
