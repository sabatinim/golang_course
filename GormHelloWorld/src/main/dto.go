package main

import (
	"database/sql" //note: import from base golang packages
	"time"
)
type User struct {
	ID        int
	Birthday  time.Time
	Age       int
	Name      string `sql:"size:255"` // Default size for string is 255, you could reset it with this tag
	Num       int    `sql:"AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	Emails            []Email       // One-To-Many relationship (has many)
	BillingAddress    Address       // One-To-One relationship (has one)
	BillingAddressID  sql.NullInt64 // Foreign key of BillingAddress
	ShippingAddress   Address       // One-To-One relationship (has one)
	ShippingAddressID int           // Foreign key of ShippingAddress
	IgnoreMe          int           `sql:"-"`                          // Ignore this field
	Languages         []Language    `gorm:"many2many:user_languages;"` // Many-To-Many relationship, 'user_languages' is join table
}

type Email struct {
	ID         int
	UserID     int    `sql:"index"`                          // Foreign key (belongs to), tag `index` will create index for this field when using AutoMigrate
	Email      string `sql:"type:varchar(100);unique_index"` // Set field's sql type, tag `unique_index` will create unique index
	Subscribed bool
}

type Address struct {
	ID       int
	Address1 string         `sql:"not null;unique"` // Set field as not nullable and unique
	Address2 string         `sql:"type:varchar(100);unique"`
	//Post     sql.NullString `sql:"not null"`
}


type Language struct {
	ID   int
	Name string `sql:"index:idx_name_code"` // Create index with name, and will create combined index if find other fields defined same name
	Code string `sql:"index:idx_name_code"` // `unique_index` also works
}
