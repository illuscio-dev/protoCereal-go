package messagesCereal

import (
	"errors"
	"fmt"
	googleUUID "github.com/google/uuid"
	mongoUUID "go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

var ErrUUIDByteLen = errors.New("proto uuid message must be 16 bytes")

// Validates that a protobuf UID value is the correct number of bytes.
func (x *UUID) Validate() error {
	if len(x.Value) != 16 {
		return fmt.Errorf("%w: %v bytes found", ErrUUIDByteLen, len(x.Value))
	}
	return nil
}

// Converts cereal UUID value into google UUID value.
func (x *UUID) ToGoogle() (googleUUID.UUID, error) {
	if err := x.Validate(); err != nil {
		return [16]byte{}, err
	}
	return googleUUID.FromBytes(x.Value)
}

// Converts cereal UUID value into mongo helper UUID value (NOT the binary primitive
// value).
func (x *UUID) ToMongo() (mongoUUID.UUID, error) {
	if err := x.Validate(); err != nil {
		return [16]byte{}, err
	}
	result := mongoUUID.UUID{}
	for i, thisByte := range x.Value {
		result[i] = thisByte
	}

	return result, nil
}

// Create a new protobuf UUID value from a Google UUID value.
func UUIDFromGoogle(value googleUUID.UUID) *UUID {
	return &UUID{
		Value: value[:],
	}
}

// Create a new protobuf UUID value from a Mongo UUID value.
func UUIDFromMongo(value mongoUUID.UUID) *UUID {
	return &UUID{
		Value: value[:],
	}
}
