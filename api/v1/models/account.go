package models

type Account struct {
	AccNumber int32   `bson:"account_number"`
	Name      string  `bson:"name"`
	Balance   float64 `bson:"balance"`
}

type AccNumber struct {
	Accno int32 `protobuf:"varint,1,opt,name=accno,proto3" json:"accno,omitempty"`
}

type AccountString struct {
	AccNumber string `bson:"account_number"`
	Name      string `bson:"name"`
	Balance   string `bson:"balance"`
}
