package domain

import ticket "github.com/gomessguii/cinetria/api/ticket/models"

type ITicketRepository interface {
	Get(id int64) (*ticket.Ticket, error)
	Create(*ticket.Ticket) error
}