package store

import (
	"errors"
	"sync"
)

var (
	ErrInvalidAddress     = errors.New("wallet address is required")
	ErrInvalidBalance     = errors.New("wallet balance cannot be negative")
	ErrInvalidAmount      = errors.New("transfer amount must be positive")
	ErrWalletExists       = errors.New("wallet already exists")
	ErrWalletNotFound     = errors.New("wallet not found")
	ErrInsufficientFunds  = errors.New("insufficient funds")
	ErrSameWalletTransfer = errors.New("cannot transfer to the same wallet")
)

type Wallet struct {
	Address string
	Balance int64
}

type WalletStore struct {
	mu      sync.RWMutex
	wallets map[string]int64
}

func NewWalletStore() *WalletStore {
	return &WalletStore{
		wallets: make(map[string]int64),
	}
}

func (s *WalletStore) Create(address string, balance int64) (Wallet, error) {
	if address == "" {
		return Wallet{}, ErrInvalidAddress
	}
	if balance < 0 {
		return Wallet{}, ErrInvalidBalance
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.wallets[address]; ok {
		return Wallet{}, ErrWalletExists
	}

	s.wallets[address] = balance
	return Wallet{
		Address: address,
		Balance: balance,
	}, nil
}

func (s *WalletStore) Get(address string) (Wallet, error) {
	if address == "" {
		return Wallet{}, ErrInvalidAddress
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	balance, ok := s.wallets[address]
	if !ok {
		return Wallet{}, ErrWalletNotFound
	}

	return Wallet{
		Address: address,
		Balance: balance,
	}, nil
}

func (s *WalletStore) Transfer(from, to string, amount int64) (Wallet, Wallet, error) {
	if from == "" || to == "" {
		return Wallet{}, Wallet{}, ErrInvalidAddress
	}
	if from == to {
		return Wallet{}, Wallet{}, ErrSameWalletTransfer
	}
	if amount <= 0 {
		return Wallet{}, Wallet{}, ErrInvalidAmount
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	fromBalance, ok := s.wallets[from]
	if !ok {
		return Wallet{}, Wallet{}, ErrWalletNotFound
	}

	toBalance, ok := s.wallets[to]
	if !ok {
		return Wallet{}, Wallet{}, ErrWalletNotFound
	}

	if fromBalance < amount {
		return Wallet{}, Wallet{}, ErrInsufficientFunds
	}

	fromBalance -= amount
	toBalance += amount

	s.wallets[from] = fromBalance
	s.wallets[to] = toBalance

	return Wallet{
			Address: from,
			Balance: fromBalance,
		}, Wallet{
			Address: to,
			Balance: toBalance,
		}, nil
}
