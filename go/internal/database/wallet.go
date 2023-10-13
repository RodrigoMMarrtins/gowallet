package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Wallet struct {
	db      *sql.DB
	ID      string
	Balance float32
	//TransactionHistory []string
	UserID string
}

func NewWallet(db *sql.DB) *Wallet {
	return &Wallet{db: db}
}

func (w *Wallet) Create(userID string, balance float32) (*Wallet, error) {
	id := uuid.New().String()
	query := `INSERT INTO wallets (id, balance, user_id) VALUES ($1, $2, $3)`
	_, err := w.db.Exec(query, id, balance, userID)
	if err != nil {
		return nil, err
	}

	return &Wallet{
		ID:      id,
		Balance: balance,
		UserID:  userID,
	}, nil
}
