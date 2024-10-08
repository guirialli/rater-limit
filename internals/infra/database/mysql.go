package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/guirialli/rater_limit/config"
	"github.com/guirialli/rater_limit/pkg/utils"
	"strconv"
)

type MySql struct {
	hostname string
	port     int
	user     string
	password string
	database string
}

func NewMySql() (*MySql, error) {
	dbCfg, err := config.LoadDatabaseConfig()
	if err != nil {
		return nil, err
	}

	port, err := strconv.Atoi(dbCfg.Port)
	if err != nil {
		return nil, err
	}

	return &MySql{
		user:     dbCfg.User,
		port:     port,
		password: dbCfg.Password,
		hostname: dbCfg.Hostname,
		database: dbCfg.Database,
	}, nil
}

func (d *MySql) GetConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		d.user, d.password, d.hostname, d.port, d.database)
}

func (d *MySql) GetConnection() (*sql.DB, error) {
	dns := d.GetConnectionString()
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (d *MySql) TryConnection() error {
	_, err := d.GetConnection()
	return err
}
func (d *MySql) InitDatabase(file string) error {
	db, err := d.GetConnection()
	if err != nil {
		return err
	}

	if _, err := db.Exec("CREATE DATABASE IF NOT EXISTS " + d.database); err != nil {
		return err
	}
	if err := utils.NewDatabaseUtils().ExecScript(db, file); err != nil {
		return err
	}

	fmt.Println("Database Initialized with success!")
	return db.Close()
}
func (d *MySql) Migrate() error {
	return nil
}
