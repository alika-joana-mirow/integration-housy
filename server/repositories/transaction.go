package repositories

import (
	"be/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindOrders() ([]models.Transaction, error)
	GetOrderByUserId(ID int) ([]models.Transaction, error)
	AddTransaction(transactions models.Transaction) (models.Transaction, error)
	DeleteOrder(transactions models.Transaction) (models.Transaction, error)

	GetOrder(ID int) (models.Transaction, error)
	GetOneTransaction(ID string) (models.Transaction, error)
	UpdateTransaction(status string, ID string) error
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindOrders() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("Housy.City").Preload("Housy.User").Preload("User").Find(&transactions).Error

	return transactions, err
}

func (r *repository) GetOrder(ID int) (models.Transaction, error) {
	var transaction models.Transaction //baru tak tmbhin housy.user
	err := r.db.Preload("Housy.City").Preload("Housy.User").Preload("User").First(&transaction, "id = ?", ID).Error

	return transaction, err
}

func (r *repository) GetOneTransaction(ID string) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Housy").Preload("Housy.User").Preload("User").First(&transaction, "id = ?", ID).Error

	return transaction, err
}

func (r *repository) GetOrderByUserId(ID int) ([]models.Transaction, error) {
	var orderUser []models.Transaction
	err := r.db.Preload("Housy").Preload("User").Where(&orderUser, "user_id = ?", ID).Error

	return orderUser, err
}

// update status
func (r *repository) UpdateTransaction(status string, ID string) error {
	var transaction models.Transaction
	r.db.Preload("Housy").First(&transaction, ID)

	if status != transaction.Status && status == "success" {
		var housy models.Housy
		r.db.First(&housy, transaction.Housy.ID)
		// housy.Qty = Housy.Qty - 1
		r.db.Save(&housy)
	}

	transaction.Status = status
	err := r.db.Save(&transaction).Error

	return err
}

func (r *repository) AddTransaction(transactions models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("Housy.City").Create(&transactions).Error

	return transactions, err
}

func (r *repository) DeleteOrder(transactions models.Transaction) (models.Transaction, error) {
	err := r.db.Delete(&transactions).Error

	return transactions, err
}
