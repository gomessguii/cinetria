package models

import seat "github.com/gomessguii/cinetria/api/seat/models"

type Builder interface {
	WithId(id int64) Builder
	WithSeat(*seat.Seat) Builder
	WithUserId(userId int64) Builder
	WithIsPaid(isPaid bool) Builder
	Build() *Ticket
}

type builder struct {
	ticket *Ticket
}

func NewBuilder() Builder {
	return &builder{ticket: &Ticket{}}
}

func (b *builder) Build() *Ticket {
	return b.ticket
}

func (b *builder) WithId(id int64) Builder {
	b.ticket.id = id
	return b
}

func (b *builder) WithIsPaid(isPaid bool) Builder {
	b.ticket.isPaid = isPaid
	return b
}

func (b *builder) WithSeat(seat *seat.Seat) Builder {
	b.ticket.seat = seat
	return b
}

func (b *builder) WithUserId(userId int64) Builder {
	b.ticket.userID = userId
	return b
}
