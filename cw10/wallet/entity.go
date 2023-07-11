package wallet

type Wallet struct {
	ID         string `json:"id"`
	HolderName string `json:"holder_name"`
	Amount     int    `json:"amount"`
}
