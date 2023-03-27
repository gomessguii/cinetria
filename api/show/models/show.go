package models

import (
	"fmt"

	theater "github.com/gomessguii/cinetria/api/theater/models"
)

type Show struct {
	id          int64
	title       string
	description string
	theater     *theater.Theater
}

func (s *Show) String() string {
	return fmt.Sprintf("")
}

func (s *Show) IsValid() bool {
	if s == nil {
		return false
	}
	if s.title == "" {
		return false
	}
	if s.description == "" {
		return false
	}
	if s.theater == nil {
		return false
	}
	return s.theater.IsValid()
}

func (s *Show) ToRequest() *ShowRequest {
	return &ShowRequest{
		ID:          s.id,
		Title:       s.title,
		Description: s.description,
	}
}

type ShowRequest struct {
	ID          int64            `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Theater     *theater.Theater `json:"theater"`
}
