package main

import (
	"database/sql"
	"net"

	"github.com/RodrigoMMartins/gowallet/internal/database"
	"github.com/RodrigoMMartins/gowallet/internal/pb"
	"github.com/RodrigoMMartins/gowallet/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	walletDB := database.NewWallet(db)
	walletService := service.NewWalletService(*walletDB)

	grpcServer := grpc.NewServer()
	pb.RegisterWalletServiceServer(grpcServer, walletService)
	reflection.Register(grpcServer)
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
	
}

/*
import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	password := "secret"
	hash, _ := HashPassword(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)
}
*/
