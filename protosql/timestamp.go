package protosql

import (
	"database/sql"
	"database/sql/driver"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TimestampMarshaller wraps a *timestamppb.Timestamp value for marshalling to an SQL
// database.
type TimestampMarshaller struct {
	*timestamppb.Timestamp
}

// Value converts the value to a time.Time value for the sql driver to handle
func (timestamp TimestampMarshaller) Value() (driver.Value, error) {
	if timestamp.Timestamp == nil {
		return nil, nil
	}
	return timestamp.AsTime(), nil
}

// Scan converts the value from time.Time for the sql driver to handle.
func (timestamp *TimestampMarshaller) Scan(src interface{}) error {
	// We'll just piggyback of the sql.NullTime implementation
	nullTime := sql.NullTime{}
	err := nullTime.Scan(src)
	if err != nil {
		return err
	}

	// If this is a null value, return an empty marshaller.
	if !nullTime.Valid {
		*timestamp = TimestampMarshaller{}
		return nil
	}

	// Convert to *timestamppb.Timestamp
	*timestamp = TimestampMarshaller{
		Timestamp: timestamppb.New(nullTime.Time),
	}
	return nil
}

// Timestamp creates a new TimestampMarshaller for a given value
func Timestamp(value *timestamppb.Timestamp) TimestampMarshaller {
	return TimestampMarshaller{Timestamp: value}
}
