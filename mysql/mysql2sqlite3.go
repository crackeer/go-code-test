package main

import (
	"fmt"

	"flag"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	mysqlDsn     string
	sqliteDBFile string
)

var (
	mysqlDB  *gorm.DB
	sqliteDB *gorm.DB
)

func main() {
	flag.StringVar(&mysqlDsn, "mysql", "", "mysql dsn")
	flag.StringVar(&sqliteDBFile, "sqlite", "", "sqlite database file")
	flag.Parse()

	var err error
	mysqlDB, err = gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqliteDB, err = gorm.Open(sqlite.Open(sqliteDBFile), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	list := getMySQLTables()
	for _, v := range list {

		files := descTable(v)
		sql := getSqliteSQL(v, files)

		err := createSqliteTable(sql)
		if err != nil {
			panic(err)
		}
		fmt.Println("TABLE", v)
		fmt.Println("SQL:", sql)
		fmt.Println("Result:", err)
		fmt.Println("------------------------------")

	}
}

func getMySQLTables() []string {
	list := []string{}
	mysqlDB.Raw("show tables").Scan(&list)
	return list
}

type TableDesc struct {
	Field   string `json:"Field"`
	Type    string `json:"Type"`
	Key     string `json:"Key"`
	Null    string `json:"Null"`
	Default string `json:"Default"`
}

func descTable(name string) []TableDesc {
	retData := []TableDesc{}
	err := mysqlDB.Raw("desc " + name).Scan(&retData).Error
	if err != nil {
		return nil
	}
	return retData
}

func getSqliteSQL(name string, desc []TableDesc) string {
	fields := []string{}
	pri := false
	for _, v := range desc {
		v.Type = strings.Replace(v.Type, "unsigned", "", -1)
		v.Field = fmt.Sprintf("`%s`", v.Field)
		tmp := []string{
			v.Field, v.Type,
		}

		if v.Key == "PRI" && !pri {
			tmp = append(tmp, "PRIMARY KEY")
			pri = true
		}
		if v.Null == "NO" {
			tmp = append(tmp, "NOT NULL")
		}

		if v.Default == "NULL" {
			tmp = append(tmp, "DEFAULT NULL")
		} else if v.Default == "CURRENT_TIMESTAMP" {
			tmp = append(tmp, "DEFAULT CURRENT_TIMESTAMP")
		} else if v.Default != "" {
			tmp = append(tmp, fmt.Sprintf("DEFAULT '%s'", v.Default))
		}
		fields = append(fields, strings.Join(tmp, " "))
	}

	return fmt.Sprintf(" CREATE TABLE `%s` (%s)", name, strings.Join(fields, ","))
}

func createSqliteTable(sql string) error {
	return sqliteDB.Exec(sql).Error
}

func moveData(table string) {
	list := []map[string]interface{}{}
	mysqlDB.Table(table).Order("id desc").Limit(100).Find(&list)

	sqliteDB.Table(table).Create(&list)
}
