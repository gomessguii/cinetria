package models

type Builder interface {
	WithId(id int64) Builder
	WithName(name string) Builder
	WithAddress(address string) Builder
	Build() *Theater
}

type builder struct {
	theater *Theater
}

func NewBuilder() Builder {
	return &builder{&Theater{}}
}

func (b *builder) WithAddress(address string) Builder {
	b.theater.address = address
	return b
}

func (b *builder) WithId(id int64) Builder {
	b.theater.id = id
	return b
}

func (b *builder) WithName(name string) Builder {
	b.theater.name = name
	return b
}

func (b *builder) Build() *Theater {
	return b.theater
}
