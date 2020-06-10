package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var (
	// ErrNotFound error when record not found
	ErrNotFound = fmt.Errorf("record Not Found")

	// ErrUpdateFailed error when update fails
	ErrUpdateFailed = fmt.Errorf("db update error")

	// ErrInsertFailed error when insert fails
	ErrInsertFailed = fmt.Errorf("db insert error")

	// ErrDeleteFailed error when delete fails
	ErrDeleteFailed = fmt.Errorf("db delete error")

	// ErrBadParams error when bad params passed in
	ErrBadParams = fmt.Errorf("bad params error")

	// DB reference to database
	DB *gorm.DB
)
