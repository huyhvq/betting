package repository

import (
	"github.com/huyhvq/betting/pkg/model"
	"gorm.io/gorm"
	"time"
)

type WagerRepository interface {
	Create(request model.Wager) (model.Wager, error)
}

type wagerRepo struct {
	db *gorm.DB
}

func (r *wagerRepo) Create(w model.Wager) (model.Wager, error) {
	w.PlacedAt = time.Now()
	if err := r.db.Create(&w).Error; err != nil {
		return w, err
	}
	return w, nil
}

func NewWager(db *gorm.DB) WagerRepository {
	return &wagerRepo{db: db}
}
