package entity

type Account struct {
	ID       string  `gorm:"primaryKey;column:id;type:char(36);not null"`
	Name     string  `gorm:"column:name;type:varchar(100);not null"`
	Email    *string `gorm:"column:email;type:varchar(100)"`
	Password string  `gorm:"column:password;type:tinytext;not null"`
	Role     string  `gorm:"column:role;type:varchar(100);not null"`
}

func (Account) TableName() string {
	return "account"
}

var AccountColumns = struct {
	Name string
}{
	Name: "name",
}
