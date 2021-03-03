package repository

import (
	"github.com/huyhvq/betting/pkg/model"
	"gorm.io/gorm"
)

type WagerRepository interface {
	Create(request model.Wager) (*model.Wager, error)
	GetByID(id uint) (*model.Wager, error)
	Buy(id uint, price float64) (*model.Transaction, error)
	List(offset int, size int) ([]model.Wager, error)
}

type wagerRepo struct {
	db *gorm.DB
}

func (r *wagerRepo) Create(w model.Wager) (*model.Wager, error) {
	if err := r.db.Create(&w).Error; err != nil {
		return nil, err
	}
	return &w, nil
}

func (r *wagerRepo) GetByID(id uint) (*model.Wager, error) {
	w := model.Wager{
		Model: gorm.Model{ID: id},
	}
	if err := r.db.First(&w).Error; err != nil {
		return nil, err
	}
	return &w, nil
}

func (r *wagerRepo) Buy(id uint, price float64) (*model.Transaction, error) {
	w, err := r.GetByID(id)
	if err != nil {
		return nil, err
	}
	t := model.Transaction{
		WagerID:     id,
		BuyingPrice: price,
	}
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&t).Error; err != nil {
			return err
		}
		w.CurrentSellingPrice = w.CurrentSellingPrice - price
		w.AmountSold = w.SellingPrice - w.CurrentSellingPrice
		w.PercentageSold = int64((w.AmountSold / w.SellingPrice) * 100)
		if err := tx.Save(w).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *wagerRepo) List(offset int, size int) ([]model.Wager, error) {
	var ws []model.Wager
	paginate := func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(size)
	}
	if err := r.db.Scopes(paginate).Find(&ws).Error; err != nil {
		return nil, err
	}
	return ws, nil
}

func NewWager(db *gorm.DB) WagerRepository {
	return &wagerRepo{db: db}
}
