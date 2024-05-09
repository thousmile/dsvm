package initialize

import (
	"database/sql"
	"dsvm/conf"
	"dsvm/global"
	"dsvm/model"
	"fmt"
	_ "github.com/glebarez/go-sqlite"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/microsoft/go-mssqldb"
	"github.com/pressly/goose/v3"
	"log"
)

func DataSource() {
	sources := global.AppConfig.DataSources
	m1 := make(map[string]*model.GooseDataSource, len(sources))
	for k, v := range sources {
		db := getDbByDialect(v)
		err := db.Ping()
		if err != nil {
			log.Panic(err)
		}
		m1[k] = &model.GooseDataSource{
			Dialect:         v.Dialect,
			CreateSchemaSql: v.CreateSchemaSql,
			SwitchSchemaSql: v.SwitchSchemaSql,
			Db:              db,
		}
	}
	global.DataSources = m1
}

func getDbByDialect(v conf.DataSourceConfig) *sql.DB {
	var err error
	var db *sql.DB
	switch goose.Dialect(v.Dialect) {
	case goose.DialectMySQL:
		db, err = mysql(v)
		break
	case goose.DialectPostgres:
		db, err = postgres(v)
		break
	case goose.DialectMSSQL:
		db, err = mssql(v)
		break
	default:
		db, err = mysql(v)
		break
	}
	if err != nil {
		log.Println(err)
	}
	if err = db.Ping(); err != nil {
		log.Println(err)
		return nil
	}
	return db
}

func mysql(v conf.DataSourceConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		v.Username, v.Password, v.Addr, v.Port, v.DbName)
	return sql.Open("mysql", dsn)
}

func postgres(v conf.DataSourceConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable",
		v.Addr, v.Port, v.Username, v.Password, v.DbName)
	return sql.Open("postgres", dsn)
}

func mssql(v conf.DataSourceConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		v.Username, v.Password, v.Addr, v.Port, v.DbName)
	return sql.Open("mssql", dsn)
}

func sqlite3(v conf.DataSourceConfig) (*sql.DB, error) {
	return sql.Open("sqlite", v.Addr)
}
