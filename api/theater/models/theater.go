package models

import "fmt"

type Theater struct {
	id      int64
	name    string
	address string
}

func (t *Theater) IsValid() bool {
	if t == nil {
		return false
	}
	if t.name == "" {
		return false
	}
	if  t.address == "" {
		return false
	}
	return true
}

func (s *Theater) String() string {
	return fmt.Sprintf("%s - %s", s.name, s.address)
}

func (s *Theater)ToRequest() *TheaterRequest {
	return &TheaterRequest{
		ID: s.id,
		Name: s.name,
		Address: s.address,
	}
}

type TheaterRequest struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
