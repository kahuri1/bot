package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type PostgresSQL struct {
	db *sql.DB
}
type Parsel struct {
	id   int
	name string
}

func NewDBConnect() (*PostgresSQL, error) {

	viper.SetConfigName("docker-compose") // name of config file (without extension)
	viper.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	user := viper.GetString("services.db_bot.environment.POSTGRES_USER")
	password := viper.GetString("services.db_bot.environment.POSTGRES_PASSWORD")
	dbname := viper.GetString("services.db_bot.container_name.db_bot")

	dbName := fmt.Sprintf(`user=$s password=$s dbname=$s sslmode=require`, user, password, dbname)
	db, err := sql.Open("postgres", dbName)
	if err != nil {
		return &PostgresSQL{}, err
	}
	return &PostgresSQL{db}, nil
}

func (s PostgresSQL) GetTable() ([]Parsel, error) {
	row, err := s.db.Query("SELECT id, name FROM table_name")
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var res []Parsel
	for row.Next() {
		if err = row.Err(); err != nil {
			return nil, err
		}
		p := Parsel{}
		err := row.Scan(&p.id, &p.name)
		res = append(res, p)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}
