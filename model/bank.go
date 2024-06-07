package model

type Bank struct {
	BankCode string ` gorm:"primaryKey" `
	Name     string ` gorm:"column:name" `
	Address  string
}

func (a *Bank) TableName() string {
	return "bank"
}
