package goservice

import (
	"database/sql"
)

// Services contains methods for the Actions
type Services interface {
	GetFunction(ID int) (string, error)
	PostFunction(payload TestPayload) (int, error)
}

// Service implementes these methods
type Service struct {
	Db *sql.DB `inject:""`
}

// compile-time interface implementation check
var _ Services = &Service{}

// GetFunction should be replaced with real function
func (s *Service) GetFunction(ID int) (string, error) {
	var value string
	err := s.Db.QueryRow("SELECT `value` FROM `test` WHERE id = ?", ID).Scan(&value)
	if err != nil {
		return "", err
	}

	return value, nil
}

// PostFunction should be replaced with real function
func (s *Service) PostFunction(payload TestPayload) (int, error) {
	result, err := s.Db.Exec("INSERT INTO `test` (`value`) VALUES (?)", payload.Value)
	if err != nil {
		return 0, err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(lastInsertID), nil
}
