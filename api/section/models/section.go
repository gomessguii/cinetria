package models

import (
	"fmt"
	"time"

	show "github.com/gomessguii/cinetria/api/show/models"
)

type Section struct {
	id        int64
	room      string
	show      *show.Show
	startTime time.Time
}

func (s *Section) String() string {
	return fmt.Sprintf("Room: %s\n%s", s.room, s.show.String())
}

func (s *Section)IsValid() bool {
	if s == nil {
		return false
	}
	if s.room == "" {
		return false
	}
	if !s.startTime.After(time.Now()) {
		return false
	}
	if s.show == nil {
		return false
	}
	return s.show.IsValid()
}

func (s *Section) ToRequest() *SectionRequest {
	return &SectionRequest{ID: s.id, Room: s.room, Show: s.show, StartTime: s.startTime}
}

type SectionRequest struct {
	ID        int64            `json:"id"`
	Room      string           `json:"room"`
	Show      *show.Show       `json:"show"`
	StartTime time.Time        `json:"start_time"`
}
