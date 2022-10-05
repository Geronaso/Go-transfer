package datastruct

type Transfer struct {
	Id                     int64
	Account_origin_id      int64
	Account_destination_id int64
	Amount                 float64
	Created_at             string
}

type TransferValues struct {
	Origin      Account
	Destination Account
}
