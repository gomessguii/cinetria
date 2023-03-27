package models

import theater "github.com/gomessguii/cinetria/api/theater/models"

type Builder interface {
	WithID(id int64) Builder
	WithTitle(title string) Builder
	WithDescription(description string) Builder
	WithTheater(theater *theater.Theater) Builder
	Build() *Show
}

type builder struct {
	show *Show
}

func NewBuilder() Builder {
	return &builder{&Show{}}
}

func (b *builder) WithDescription(description string) Builder {
	b.show.description = description
	return b
}

func (b *builder) WithID(id int64) Builder {
	b.show.id = id
	return b
}

func (b *builder) WithTheater(theater *theater.Theater) Builder {
	b.show.theater = theater
	return b
}

func (b *builder) WithTitle(title string) Builder {
	b.show.title = title
	return b
}

func (b *builder) Build() *Show {
	return b.show
}