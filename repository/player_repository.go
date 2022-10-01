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

	sql := ("insert into player (id, name) values($1, $2)")

	_, err = conn.Exec(context.Background(), sql, player.Id, player.Name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "insert failed: %v\n", err)
		os.Exit(1)
	}
	return
}