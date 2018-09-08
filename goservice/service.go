package goservice

import (
	"database/sql"
)

// ServiceActions contains methods for the Actions
type ServiceActions interface {
	TestFunction() error
}

// Service implementes these methods
type Service struct {
	Db *sql.DB `inject:""`
}

// TestFunction will always return nil, replace with real function
func (s *Service) TestFunction() error {
	return nil
}
