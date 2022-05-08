package mysql 

import (
	"errors"
	"gorm.io/gorm"
)


func IsRecordNotFoundError(err error) bool {
	return err != nil && errors.Is(err, gorm.ErrRecordNotFound) 
}