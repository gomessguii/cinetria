package models

import (
	"time"

	"github.com/gomessguii/cinetria/api/show/models"
)

type Builder interface {
	WithID(id int64) Builder
	WithRoom(room string) Builder
	WithShow(show *models.Show) Builder
	WithStartTime(startTime time.Time) Builder
	Build() *Section
}
type builder struct {
	section *Section
}

func NewBuilder() Builder {
	return &builder{section: &Section{}}
}

func (b *builder) WithID(id int64) Builder {
	b.section.id = id
	return b
}

func (b *builder) WithRoom(room string) Builder {
	b.section.room = room
	return b
}

func (b *builder) WithShow(show *models.Show) Builder {
	b.section.show = show
	return b
}

func (b *builder) WithStartTime(startTime time.Time) Builder {
	b.section.startTime = startTime
	return b
}

func (b *builder) Build() *Section {
	return b.section
}
