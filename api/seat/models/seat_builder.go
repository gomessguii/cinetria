package seat

import section "github.com/gomessguii/cinetria/api/section/models"

type Builder interface {
	WithId(id int64) Builder
	WithLocation(location string) Builder
	WithSection(section *section.Section) Builder
	Build() *Seat
}

type builder struct{
	seat *Seat
}

func NewBuilder() Builder {
	return &builder{seat: &Seat{}}
}


func (b *builder) WithId(id int64) Builder {
	b.seat.id = id
	return b
}

func (b *builder) WithLocation(location string) Builder {
	b.seat.location = location
	return b
}

func (b *builder) WithSection(section *section.Section) Builder {
	b.seat.section = section
	return b
}

func (b *builder) Build() *Seat {
	return b.seat
}
