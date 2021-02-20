package cereal

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

// Converts cereal UUID value into google UUID value. Returns zero value if
// message pointer is nil.
func (x *UUID) ToGoogle() (googleUUID.UUID, error) {
	if x == nil {
		return [16]byte{}, nil
	}

	if err := x.Validate(); err != nil {
		return [16]byte{}, err
	}
	return googleUUID.FromBytes(x.Value)
}

// As .ToGoogle(), but panics on conversion error.
func (x *UUID) MustGoogle() googleUUID.UUID {
	val, err := x.ToGoogle()
	if err != nil {
		panic(fmt.Errorf(
			"could not convert UUID message to google UUID: %w", err),
		)
	}

	return val
}

// Converts cereal UUID value into mongo helper UUID value (NOT the binary primitive
// value).
func (x *UUID) ToMongo() (mongoUUID.UUID, error) {
	if x == nil {
		return [16]byte{}, nil
	}

	if err := x.Validate(); err != nil {
		return [16]byte{}, err
	}
	result := mongoUUID.UUID{}
	for i, thisByte := range x.Value {
		result[i] = thisByte
	}

	return result, nil
}

// As .ToMongo(), but panics on conversion error.
func (x *UUID) MustMongo() mongoUUID.UUID {
	val, err := x.ToMongo()
	if err != nil {
		panic(fmt.Errorf(
			"could not convert UUID message to Mongo UUID: %w", err),
		)
	}

	return val
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

// Generate a new random UUID using google's UUID implementation.
func NewUUIDRandom() (*UUID, error) {
	gUUID, err := googleUUID.NewRandom()
	if err != nil {
		return nil, err
	}

	return UUIDFromGoogle(gUUID), nil
}

// Generate a new random UUID using google's UUID implementation.
func MustUUIDRandom() *UUID {
	gUUID, err := NewUUIDRandom()
	if err != nil {
		panic(fmt.Errorf("error creating UUID: %w", err))
	}

	return gUUID
}
