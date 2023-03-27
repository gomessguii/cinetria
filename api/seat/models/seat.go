package seat

import (
	"fmt"

	section "github.com/gomessguii/cinetria/api/section/models"
)

type Seat struct {
	id       int64
	location string
	section  *section.Section
}

func (s *Seat) String() string {
	return fmt.Sprintf("Seat Location: %s\nSection: %s\n", s.location, s.section.String())
}

func (s *Seat)IsValid() bool {
	if s == nil {
		return false
	}
	if s.location == "" {
		return false
	}
	if s.section == nil {
		return false
	}
	return s.section.IsValid()
}

func (s *Seat) ToRequest() *SeatRequest {
	return &SeatRequest{
		ID:       s.id,
		Location: s.location,
		Section:  s.section,
	}
}

type SeatRequest struct {
	ID       int64           `json:"id"`
	Location string          `json:"location"`
	Section  *section.Section `json:"section"`
}
