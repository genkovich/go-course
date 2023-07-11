package wallet

type Storage struct {
	wallets []*Wallet
}

func NewStorage() *Storage {
	return &Storage{
		wallets: []*Wallet{
			{
				ID:         "1",
				HolderName: "John",
				Amount:     100,
			},
			{
				ID:         "2",
				HolderName: "Jane",
				Amount:     0,
			},
		},
	}
}

func (s *Storage) GetAll() []*Wallet {
	return s.wallets
}

func (s *Storage) GetWalletById(id string) (*Wallet, error) {
	for _, wallet := range s.wallets {
		if wallet.ID == id {
			return wallet, nil
		}
	}

	return nil, nil
}

func (s *Storage) AddWallet(wallet *Wallet) {
	s.wallets = append(s.wallets, wallet)
}

func (s *Storage) UpdateAmount(wallet *Wallet, amount int) {
	wallet.Amount = amount
}

func (s *Storage) DeleteWallet(walletId string) {
	for i, wallet := range s.wallets {
		if wallet.ID == walletId {
			s.wallets = append(s.wallets[:i], s.wallets[i+1:]...)
		}
	}
}
