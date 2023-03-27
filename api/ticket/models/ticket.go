package models

import seatModel "github.com/gomessguii/cinetria/api/seat/models"

type Ticket struct {
	id     int64
	seat   *seatModel.Seat
	userID int64
	isPaid bool
}

func (t *Ticket) IsValid() bool {
	if t == nil {
		return false
	}
	if t.userID == 0 {
		return false
	}
	return t.seat.IsValid()
}

func (t *Ticket) ToRequest() *TicketRequest {
	return &TicketRequest{ID: t.id, Seat: t.seat, UserID: t.userID, IsPaid: &t.isPaid}
}

type TicketRequest struct {
	ID     int64           `json:"id"`
	Seat   *seatModel.Seat `json:"seat"`
	UserID int64           `json:"user_id"`
	IsPaid *bool           `json:"is_paid"`
}
