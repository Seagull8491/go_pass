package util

import (
	"go_pass/model"
	"path"
	"testing"
)

const SQL_FILE_NAME_FOR_TEST = "test.sql"

func TestDBOpen(t *testing.T) {
	tempDir := t.TempDir()
	sql_file := path.Join(tempDir, SQL_FILE_NAME_FOR_TEST)
	_, err := DBOpen(sql_file)
	if err != nil {
		t.Fatalf("DB open failed")
	}
}

func TestCreateTable(t *testing.T) {
	tempDir := t.TempDir()
	sql_file := path.Join(tempDir, SQL_FILE_NAME_FOR_TEST)
	db, err := DBOpen(sql_file)
	if err != nil {
		t.Fatalf("DB open failed")
	}
	defer db.Close()

	t.Run("Create Table", func(t *testing.T) {
		err := CreateTable(db)
		if err != nil {
			t.Fatalf("CreateTable is failed: %s", err)
		}
	})
}

func TestInsertPassRecord(t *testing.T) {
	tempDir := t.TempDir()
	sql_file := path.Join(tempDir, SQL_FILE_NAME_FOR_TEST)
	db, err := DBOpen(sql_file)
	if err != nil {
		t.Fatalf("DB open failed")
	}
	defer db.Close()

	err = CreateTable(db)
	if err != nil {
		t.Fatalf("CreateTable is failed: %s", err)
	}

	t.Run("Check argument validation succeeds: entity = ``", func(t *testing.T) {
		err := InsertPassRecord(db, "", "password_a")
		if err == nil {
			t.Fail()
		}
	})

	t.Run("Check argument validation succeeds: pass_encrypted = ``", func(t *testing.T) {
		err := InsertPassRecord(db, "site_a", "")
		if err == nil {
			t.Fail()
		}
	})

	t.Run("Check inserting a record succeeds", func(t *testing.T) {
		err := InsertPassRecord(db, "site_a", "password_a")
		if err != nil {
			t.Fatalf("Insert is failed")
		}
	})
}

func TestSelectPassRecordByEntity(t *testing.T) {
	tempDir := t.TempDir()
	sql_file := path.Join(tempDir, SQL_FILE_NAME_FOR_TEST)
	db, err := DBOpen(sql_file)
	if err != nil {
		t.Fatalf("DB open failed")
	}
	defer db.Close()

	err = CreateTable(db)
	if err != nil {
		t.Fatalf("CreateTable is failed: %s", err)
	}

	err = InsertPassRecord(db, "test_entity", "test_pass_encrypted")
	if err != nil {
		t.Fatalf("InsertPassRecord is failed: %s", err)
	}

	t.Run("Check argument validation succeeds: entity_name = ``", func(t *testing.T) {
		entity := model.PassEntity{
			Name:          "",
			PassRawText:   "",
			PassEncrypted: "",
		}
		err := SelectPassRecordByEntity(db, "", &entity)
		if err == nil {
			t.Fail()
		}
	})

	t.Run("Check selecting a record succeeds", func(t *testing.T) {
		entity := model.PassEntity{
			Name:          "",
			PassRawText:   "",
			PassEncrypted: "",
		}
		err := SelectPassRecordByEntity(db, "test_entity", &entity)
		if err != nil {
			t.Fatalf("SelectPassRecordByEntity is failed: %s", err)
		}

		want_entity := "test_entity"
		got_entity := entity.Name
		if want_entity != got_entity {
			t.Fatalf("want_entity: [%s], got_entity: [%s]", want_entity, got_entity)
		}
		want_pass_encrypted := "test_pass_encrypted"
		got_pass_encrypted := entity.PassEncrypted
		if want_pass_encrypted != got_pass_encrypted {
			t.Fatalf("want_pass_encrypted: [%s], got_pass_encrypted: [%s]", want_pass_encrypted, got_pass_encrypted)
		}
	})

}
