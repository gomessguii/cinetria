package domain

import (
	"fmt"
	"testing"
	"time"

	seatModel "github.com/gomessguii/cinetria/api/seat/models"
	sectionModel "github.com/gomessguii/cinetria/api/section/models"
	showModel "github.com/gomessguii/cinetria/api/show/models"
	theaterModel "github.com/gomessguii/cinetria/api/theater/models"
	ticketModel "github.com/gomessguii/cinetria/api/ticket/models"
	internal_errors "github.com/gomessguii/cinetria/internal/errors"
)

type repositoryMock struct {
	data map[int64]*ticketModel.Ticket
	index int64
}

func (m *repositoryMock) Get(id int64) (*ticketModel.Ticket, error) {
	found := m.data[id]
	if found == nil {
		return nil, internal_errors.NewErrorWithCode(404, fmt.Errorf("not found"))
	}
	return found, nil
}

func (m *repositoryMock) Create(ticket *ticketModel.Ticket) error {
	if !ticket.IsValid() {
		return internal_errors.NewErrorWithCode(400, fmt.Errorf("not valid"))
	}
	m.data[m.index] = ticket
	m.index++
	return nil
}

func TestReserveTicket(t *testing.T) {
	ticketDomain := New(&repositoryMock{data: make(map[int64]*ticketModel.Ticket)})

	theater := theaterModel.NewBuilder().WithName("Cine Go").WithAddress("100 Gopher St.").Build()
	show := showModel.NewBuilder().WithTitle("Star Wars: Episode IV - A New Hope").WithDescription("A long time ago in a galaxy far, far away…").WithTheater(theater).Build()
	section := sectionModel.NewBuilder().WithRoom("2").WithStartTime(time.Now().Add(time.Hour)).WithShow(show).Build()
	seat := seatModel.NewBuilder().WithLocation("A1").WithSection(section).Build()
	ticket, err := ticketDomain.ReserveTicket(seat, 1)
	if err != nil {
		t.Fatal(err)
	}
	if !ticket.IsValid() {
		t.Error("invalid ticket")
	}
}
