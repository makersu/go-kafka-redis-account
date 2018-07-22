package main

import (
	"errors"
	"strconv"
)

type Account struct {
	Id      string
	Name    string
	Balance int
}

func updateAccount(id string, data map[string]interface{}) (*Account, error) {
	cmd := Redis.HMSet(id, data)

	if err := cmd.Err(); err != nil {
		return nil, err
	} else {
		return FetchAccount(id)
	}
}

func FetchAccount(id string) (*Account, error) {
	cmd := Redis.HGetAll(id)
	if err := cmd.Err(); err != nil {
		return nil, err
	}

	data := cmd.Val()
	if len(data) == 0 {
		return nil, nil
	} else {
		return ToAccount(data)
	}
}

func ToAccount(m map[string]string) (*Account, error) {
	balance, err := strconv.Atoi(m["Balance"])
	if err != nil {
		return nil, err
	}

	if _, ok := m["Id"]; !ok {
		return nil, errors.New("Missing account ID")
	}

	return &Account{
		Id:      m["Id"],
		Name:    m["Name"],
		Balance: balance,
	}, nil
}
