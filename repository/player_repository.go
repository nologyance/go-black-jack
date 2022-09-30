package repository

import (
	"black-jack/entity"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

type PlayerRepository interface {
	InsertPlayer(player entity.PlayerEntity) (id int, err error)
}

type playerRepository struct {}

func NewPlayerRepository() PlayerRepository {
	return &playerRepository{}
}

func (pr *playerRepository) InsertPlayer(player entity.PlayerEntity) (id int, err error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	var i int
	var name string

	err = conn.QueryRow(context.Background(), "select * from player;").Scan(&i, &name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(i)
	fmt.Println(name)
	return
}