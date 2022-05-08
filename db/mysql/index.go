package mysql

import (
	"fmt"
	"log"

	"sturm/vars"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func InitDB() {
	DB = Connect()
}

// Connect to database
// https://github.com/go-sql-driver/mysql/issues/461 about connection keep alive or open/close it
func Connect() *gorm.DB {
	connectionString := getConnectionString()

	// dbConnection, connectionString := getConnectionString()
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		CreateBatchSize: 100000,
		Logger: logger.Default.LogMode(logger.Silent),
	})
    
	// Для создания и обновления с вложенными моделями
	db = db.Session(&gorm.Session{FullSaveAssociations: true})           

	if err != nil {
		log.Fatalf("DB error: %s", err.Error())
	}

	// dbConfig, _ := db.DB()
	// dbConfig.SetMaxOpenConns(250)
	// dbConfig.SetMaxIdleConns(1000)
	// dbConfig.SetConnMaxLifetime(time.Minute * 86000)  // time.Hour
    
	return db
}              

// Ping to database
func Ping(){
	db := Connect()
	Disconnect(db)
} 

// Отключение базы данных
func Disconnect(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err!=nil{
	   fmt.Println("Ошибка отключения базы данных.", err.Error())
	}
	sqlDB.Close()
}

// Get Connect string
func getConnectionString() (string) {
	// dbConnection := vars.DB_CONNECTION
	dbUser := vars.Config.DB_USER
	dbPwd  := vars.Config.DB_PASSWORD
	dbName := vars.Config.DB_DATABASE
	dbHost := vars.Config.DB_HOST
	return fmt.Sprintf(
		"%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local&timeout=100s", 
		dbUser, dbPwd, dbHost, dbName,
	)
}