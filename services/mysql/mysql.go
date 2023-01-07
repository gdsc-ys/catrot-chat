package sql

import (
	"database/sql"
	"log"
	"os"
	"sync"
	"time"

	"github.com/go-sql-driver/mysql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type dbConnectionPool struct {
	Write *sql.DB
	Read  *sql.DB
}

var MasterBun *bun.DB
var SlaveBun *bun.DB

func SetDBConnection(wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	apiEnv := os.Getenv("API_ENV")
	var mySqlConnection *dbConnectionPool
	if apiEnv == "production" {
		log.Println("프로덕션 디비 연결...")
		mySqlConnection = setProductionDBConnection()
	} else {
		log.Println("개발 디비 연결...")
		mySqlConnection = setDevDBConnection()
	}

	errWrite := mySqlConnection.Write.Ping()
	errRead := mySqlConnection.Read.Ping()

	if errWrite != nil {
		panic(errWrite)
	}
	if errRead != nil {
		panic(errRead)
	}

	MasterBun = bun.NewDB(mySqlConnection.Write, mysqldialect.New())
	SlaveBun = bun.NewDB(mySqlConnection.Read, mysqldialect.New())
}

func setDevDBConnection() *dbConnectionPool {

	config := mysql.Config{
		User:                 "root",
		Passwd:               "root123!",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		AllowNativePasswords: true,
	}

	db, err := otelsql.Open("mysql", config.FormatDSN(),
		otelsql.WithAttributes(semconv.DBSystemMySQL),
		otelsql.WithDBName("DevDB"))
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	log.Println(string("\033[32m"), "개발 디비 커넥션 성공", string("\033[0m"))

	dbConnection := dbConnectionPool{}
	dbConnection.Write = db
	dbConnection.Read = db

	return &dbConnection
}

func setProductionDBConnection() *dbConnectionPool {

	configWrite := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PW"),
		Net:                  os.Getenv("DB_NET"),
		Addr:                 os.Getenv("DB_ADDR"),
		AllowNativePasswords: true,
	}

	configRead := mysql.Config{
		User: 				  os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PW"),
		Net:                  os.Getenv("DB_NET"),
		Addr:                 os.Getenv("DB_ADDR"),
		AllowNativePasswords: true,
	}

	dbWrite, errWrite := otelsql.Open("mysql", configWrite.FormatDSN(),
		otelsql.WithAttributes(semconv.DBSystemMySQL),
		otelsql.WithDBName("RealMasterDB"))
	if errWrite != nil {
		panic(errWrite)
	}

	dbWrite.SetConnMaxLifetime(time.Minute * 3)
	dbWrite.SetMaxOpenConns(150)
	dbWrite.SetMaxIdleConns(150)
	log.Println(string("\033[32m"), "쓰기 디비 커넥션 성공", string("\033[0m"))

	dbRead, errRead := otelsql.Open("mysql", configRead.FormatDSN(),
		otelsql.WithAttributes(semconv.DBSystemMySQL),
		otelsql.WithDBName("RealSlaveDB"))
	if errRead != nil {
		panic(errRead)
	}

	dbRead.SetConnMaxIdleTime(time.Minute * 3)
	dbRead.SetMaxOpenConns(150)
	dbRead.SetMaxIdleConns(150)
	log.Println(string("\033[32m"), "읽기 디비 커넥션 성공", string("\033[0m"))

	dbConnection := dbConnectionPool{}
	dbConnection.Write = dbWrite
	dbConnection.Read = dbRead

	return &dbConnection
}
