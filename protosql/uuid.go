package protosql

import (
	"database/sql/driver"
	"fmt"
	"github.com/illuscio-dev/protoCereal-go/cereal"
)

// UUIDHexMarshaller wraps *cereal.UUID and marshals to a hex-string format.
type UUIDHexMarshaller struct {
	*cereal.UUID
}

// Value converts the value to a time.Time value for the sql driver to handle
func (uuid UUIDHexMarshaller) Value() (driver.Value, error) {
	if uuid.UUID == nil {
		return nil, nil
	}

	googleVal, err := uuid.ToGoogle()
	if err != nil {
		return nil, fmt.Errorf("error converting uuid bytes: %w", err)
	}

	return googleVal.String(), nil
}

func (uuid *UUIDHexMarshaller) Scan(src interface{}) error {
	// If this src is nil, set the value to a
	if src == nil {
		*uuid = UUIDHexMarshaller{UUID: nil}
		return nil
	}

	if uuid.UUID == nil {
		uuid.UUID = new(cereal.UUID)
	}

	return uuid.UUID.Scan(src)
}

// UUIDHex creates a new hex string marshaller for cereal.UUID values.
func UUIDHex(uuid *cereal.UUID) UUIDHexMarshaller {
	return UUIDHexMarshaller{UUID: uuid}
}
