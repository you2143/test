package response

import (
	"vote-app-api/domain/model"
)

type MasterMovie struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
}

type MasterMovies struct {
	MasterMovies []model.MasterMovie `json:"MasterMovies"`
}
