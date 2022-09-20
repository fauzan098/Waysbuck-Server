package repositories

import (
	"bewaysbuck/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransactions(ID int) ([]models.Transaction, error)
	GetTransactionId() (models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	GetIdTransaction(ID string) (models.Transaction, error)
	GetDetailTransaction(ID int) (models.Transaction, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	UpdateTransaction(transaction models.Transaction) (models.Transaction, error)
	DeleteTransaction(transaction models.Transaction) (models.Transaction, error)
	UpdateTransactions(status string, ID string) error
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransactions(ID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("User").Preload("Cart").Preload("Cart.Product").Preload("Cart.Topping").Find(&transactions, "user_id = ?", ID).Error

	return transactions, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").Preload("Cart").Preload("Cart.Product").Preload("Cart.Topping").First(&transaction, "user_id = ?", ID).Error

	return transaction, err
}

func (r *repository) GetTransactionId() (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").Preload("Cart").Preload("Cart.Product").Preload("Cart.Topping").First(&transaction, "status = ?", "Active").Error

	return transaction, err
}

func (r *repository) GetIdTransaction(ID string) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").Preload("Cart").Preload("Cart.Product").Preload("Cart.Topping").First(&transaction, "id = ?", ID).Error

	return transaction, err
}

func (r *repository) GetDetailTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").Preload("Cart").Preload("Cart.Product").Preload("Cart.Topping").First(&transaction, ID).Error

	return transaction, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error

	return transaction, err
}

func (r *repository) UpdateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Save(&transaction).Error

	return transaction, err
}

func (r *repository) DeleteTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Delete(&transaction).Error

	return transaction, err
}

func (r *repository) UpdateTransactions(status string, ID string) error {
	var transaction models.Transaction
	r.db.Preload("Product").First(&transaction, ID)

	if status != transaction.Status && status == "success" {
		var product models.Product
		r.db.First(&product, transaction.ID)
	}

	transaction.Status = status

	err := r.db.Save(&transaction).Error

	return err
}
