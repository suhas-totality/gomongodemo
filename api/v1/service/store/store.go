package store

type Store struct {
	Account Account
}

func NewStore() *Store {
	return &Store{
		Account: *NewAccount(),
	}
}
