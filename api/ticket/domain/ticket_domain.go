package domain

import (
	"fmt"

	seat "github.com/gomessguii/cinetria/api/seat/models"
	"github.com/gomessguii/cinetria/api/ticket/models"
	ticket "github.com/gomessguii/cinetria/api/ticket/models"
)

type ITicketDomain interface {
	ReserveTicket(seat *seat.Seat, userId int64) (*ticket.Ticket, error)
}

type ticketDomain struct {
	repo ITicketRepository
}

func New(repo ITicketRepository) ITicketDomain {
	return &ticketDomain{repo: repo}
}

func (d *ticketDomain) ReserveTicket(seat *seat.Seat, userId int64) (*ticket.Ticket, error) {
	if !seat.IsValid() {
		return nil, fmt.Errorf("seat is not valid")
	}

	ticket := models.NewBuilder().WithSeat(seat).WithUserId(userId).Build()
	if !ticket.IsValid() {
		return nil, fmt.Errorf("ticket is not valid")
	}

	err := d.repo.Create(ticket)
	if err != nil {
		return nil, err
	}

	return ticket, nil
}