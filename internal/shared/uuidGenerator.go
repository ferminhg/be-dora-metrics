package shared

import "github.com/google/uuid"

type UUIDGenerator interface {
	New() uuid.UUID
	Parse(s string) (uuid.UUID, error)
}

type GoogleUUIDGenerator struct{}

func (r *GoogleUUIDGenerator) New() uuid.UUID {
	return uuid.New()
}

func (r *GoogleUUIDGenerator) Parse(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}
