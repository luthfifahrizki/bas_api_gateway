package model

import "time"

type Transaction struct {
	ID              int    ` gorm:"primaryKey" `
	AccountID       string ` gorm:"foreignKey" `
	BankID          string ` gorm:"fpreignKey" `
	Amount          int    ` gorm:"column:amount" `
	TransactionDate *time.Time
}

func (a *Transaction) TableName() string {
	return "transaction"
}
