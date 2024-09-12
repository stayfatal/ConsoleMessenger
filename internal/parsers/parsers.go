package parsers

import (
	"errors"

	"github.com/google/uuid"
)

func FromAnyToUUID(obj any) (uuid.UUID, error) {
	unparsed, ok := obj.(string)
	if !ok {
		return uuid.Nil, errors.New("cant convert to string")
	}

	parsed, err := uuid.Parse(unparsed)
	if err != nil {
		return uuid.Nil, err
	}
	return parsed, nil
}
