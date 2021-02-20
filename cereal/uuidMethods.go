package cereal

import (
	"database/sql/driver"
	"errors"
	"fmt"
	googleUUID "github.com/google/uuid"
	mongoUUID "go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

var ErrUUIDByteLen = errors.New("proto uuid message must be 16 bytes")

// Validates that a protobuf UID value is the correct number of bytes.
func (x *UUID) Validate() error {
	if len(x.Bin) != 16 {
		return fmt.Errorf("%w: %v bytes found", ErrUUIDByteLen, len(x.Bin))
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
	return googleUUID.FromBytes(x.Bin)
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
	for i, thisByte := range x.Bin {
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

// Scan unmarshalls a value from an SQL database
//
// src can be string or byte.
func (x *UUID) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	// If this is a byte slice, cast to UUID and return validation.
	if bin, ok := src.([]byte); ok {
		// We need to copy the bytes here since the underlying memory is owned by the
		// driver.
		binCopy := make([]byte, len(bin))
		copy(binCopy, bin)
		*x = UUID{Bin: binCopy}
		return x.Validate()
	}

	str, ok := src.(string)
	if !ok {
		return errors.New("*cereal.UUID must be source type []byte or string")
	}

	googleVal, err := googleUUID.Parse(str)
	if err != nil {
		return fmt.Errorf("error parsing uuid string: %w", err)
	}

	*x = UUID{Bin: googleVal[:]}

	return nil
}

// Value stores the data as a binary blob. For string values, use one of the wrapper
// types in protoSQL
func (x *UUID) Value() (driver.Value, error) {
	return x.Bin, nil
}

// Create a new protobuf UUID value from a Google UUID value.
func UUIDFromGoogle(value googleUUID.UUID) *UUID {
	return &UUID{
		Bin: value[:],
	}
}

// Create a new protobuf UUID value from a Mongo UUID value.
func UUIDFromMongo(value mongoUUID.UUID) *UUID {
	return &UUID{
		Bin: value[:],
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
