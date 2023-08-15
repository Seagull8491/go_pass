package model

import (
	"errors"
	"go_pass/util"
)

type PassEntity struct {
	name          string
	passRawText   string
	passEncrypted string
}

func New(entityName string, password string, encryptKey []byte) (PassEntity, error) {
	passEncrypted, err := util.Encrypt(password, encryptKey)
	if err != nil {
		return PassEntity{}, err
	}
	return PassEntity{
		name:          entityName,
		passRawText:   password,
		passEncrypted: passEncrypted,
	}, nil
}

func (entry *PassEntity) GetPassEntryString() (string, error) {
	if entry.name != "" && entry.passEncrypted != "" {
		return entry.name + " " + entry.passEncrypted, nil
	} else {
		return "", errors.New("Invalid entry or password")
	}
}
