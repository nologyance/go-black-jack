package repository

import (
	"black-jack/entity"
	"fmt"
)

type PlayerRepository interface {
	InsertPlayer(player entity.PlayerEntity) (id int, err error)
}

type playerRepository struct {}

func NewPlayerRepository() PlayerRepository {
	return &playerRepository{}
}

func (pr *playerRepository) InsertPlayer(player entity.PlayerEntity) (id int, err error) {
	fmt.Print(player)
	return
}