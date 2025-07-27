package example

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name  *string
	Age   uint8
	Grade string
}

func strPtr(s string) *string {
	return &s
}

// Account 账户模型
type Account struct {
	ID      uint
	Balance decimal.Decimal `gorm:"type:decimal(19,4);not null"` // 高精度金额字段
}

// Transaction 转账记录模型
type Transaction struct {
	ID            uint
	FromAccountID uint            `gorm:"not null"`
	ToAccountID   uint            `gorm:"not null"`
	Amount        decimal.Decimal `gorm:"type:decimal(19,4);not null"` // 高精度金额字段
}

func Run(db *gorm.DB) {
	//student := Student{Name: strPtr("张三"), Age: 20, Grade: "三年级"}
	//result := db.Create(&student)
	//fmt.Println(result.RowsAffected)
	//fmt.Println(student)

	//var student Student
	//db.Find(&student, "age > ?", 18)
	//fmt.Println(student)

	//db.Model(&Student{}).Where("name = ?", "张三").Updates(map[string]interface{}{"grade": "四年级"})

	//result := db.Debug().Delete(&Student{}, "age < ?", 15)
	//fmt.Println(result.RowsAffected)

	//var student Student
	//db.Unscoped().Last(&student)
	//fmt.Println(student)

	a := Account{ID: 1, Balance: decimal.NewFromFloat(100.00)}
	b := Account{ID: 2, Balance: decimal.NewFromFloat(50.00)}
	db.Create(&a)
	db.Create(&b)

	var transferMoney decimal.Decimal = decimal.NewFromFloat(100)
	err := db.Transaction(func(tx *gorm.DB) error {
		if transferMoney.LessThanOrEqual(decimal.Zero) {
			return errors.New("转账金额必须大于零")
		}
		if a.Balance.LessThan(transferMoney) {
			return fmt.Errorf("账户余额不足（当前余额: %s)", a.Balance.String())
		}
		a.Balance = a.Balance.Sub(transferMoney)
		if err := tx.Save(&a).Error; err != nil {
			return fmt.Errorf("更新转出账户失败")
		}
		b.Balance = b.Balance.Add(transferMoney)
		if err := tx.Save(&b).Error; err != nil {
			return fmt.Errorf("更新转入账户失败")
		}
		transaction := Transaction{
			FromAccountID: a.ID,
			ToAccountID:   b.ID,
			Amount:        transferMoney,
		}
		if err := tx.Save(&transaction).Error; err != nil {
			return fmt.Errorf("记录转账日志失败")
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

}
