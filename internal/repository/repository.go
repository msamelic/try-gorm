package repository

import (
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"try-gorm/internal/model"
)

var _ model.Producer = (*Repo)(nil)

type Repo struct {
	db  *gorm.DB
	log *zap.Logger
}

func New(dsn string, log *zap.Logger) (model.Producer, error) {
	db, err := NewDb(dsn)
	if err != nil {
		return nil, err
	}

	return &Repo{db, log}, nil
}

func NewDb(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&model.Product{}); err != nil {
		return nil, err
	}

	return db, nil
}

func (r *Repo) List() ([]model.Product, error) {
	var products []model.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *Repo) Get(id int) (model.Product, error) {
	var product model.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (r *Repo) Create(product model.Product) (model.Product, error) {
	if err := r.db.Create(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (r *Repo) Update(product model.Product) (model.Product, error) {
	if err := r.db.Updates(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (r *Repo) Delete(product model.Product) error {
	if err := r.db.Delete(&product).Error; err != nil {
		return err
	}

	return nil
}
