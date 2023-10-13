package service

import (
	"context"

	"github.com/RodrigoMMartins/gowallet/internal/database"
	"github.com/RodrigoMMartins/gowallet/internal/pb"
)

type WalletService struct {
	pb.UnimplementedWalletServiceServer
	WalletDB database.Wallet
}

func NewWalletService(walletDB database.Wallet) *WalletService {
	return &WalletService{
		WalletDB: walletDB,
	}
}

func (w *WalletService) CreateWallet(ctx context.Context, in *pb.CreateWalletRequest) (*pb.Wallet, error) {
	wallet, err := w.WalletDB.Create(in.UserId, in.Balance)
	if err != nil {
		return nil, err
	}

	WalletResponse := &pb.Wallet{
		Id: wallet.ID,
		Balance: wallet.Balance,
		UserId: wallet.UserID,
	}

	return WalletResponse, nil
}