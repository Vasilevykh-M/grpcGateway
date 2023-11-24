package config

import (
	"encoding/json"
	"io"
	"os"
)

type ConnStructDB struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

func ReadConfDBConn(path string) (*ConnStructDB, error) {
	jsonFile, err := os.Open(path)
	defer jsonFile.Close()
	if err != nil {
		return nil, err
	}

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var connStr ConnStructDB
	err = json.Unmarshal(byteValue, &connStr)

	if err != nil {
		return nil, err
	}

	return &connStr, nil
}
