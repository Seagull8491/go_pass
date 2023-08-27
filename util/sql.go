package util

import (
	"database/sql"
	"errors"
	"fmt"
	"go_pass/model"

	_ "github.com/mattn/go-sqlite3"
)

const PASSWORD_TABLE_NAME = "password"

func DBOpen(sqliteFile string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", sqliteFile)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateTable(db *sql.DB) error {
	tableName := PASSWORD_TABLE_NAME
	cmd := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			entity TEXT NOT NULL,
			pass_encrypted TEXT NOT NULL)
	`, tableName)
	_, err := db.Exec(cmd)
	if err != nil {
		return err
	}
	return nil
}

func InsertPassRecord(db *sql.DB, entity string, pass_encrypted string) error {
	if entity == "" || pass_encrypted == "" {
		return errors.New("Neither entity or pass_encrypted can be empty.")
	}

	cmd := fmt.Sprintf("INSERT INTO %s (entity, pass_encrypted) VALUES ($1, $2)", PASSWORD_TABLE_NAME)
	_, err := db.Exec(cmd, entity, pass_encrypted)
	if err != nil {
		return err
	}
	return nil
}

func SelectPassRecordByEntity(db *sql.DB, entity_name string, pass_entity *model.PassEntity) error {
	if entity_name == "" {
		return errors.New("entity_name, pass_entity must be not empty.")
	}

	query := fmt.Sprintf("SELECT entity, pass_encrypted FROM %s WHERE entity = $1", PASSWORD_TABLE_NAME)
	var name, pass_e string
	err := db.QueryRow(query, entity_name).Scan(&name, &pass_e)
	pass_entity.Name = name
	pass_entity.PassEncrypted = pass_e
	fmt.Printf("name: [%s], pass_e:[%s]\n", name, pass_e)
	if err != nil {
		return err
	}
	return nil
}
