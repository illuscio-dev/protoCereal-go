package protosql

import (
	"database/sql/driver"
	"fmt"
	"google.golang.org/protobuf/proto"
)

// MessageBlobMarshaller marshals a proto message to a blob and back for storing
// protobufs directly to a blob field.
type MessageBlobMarshaller struct {
	proto.Message
}

// Value converts the value to a blob ([]byte) value for the sql driver to handle
func (message MessageBlobMarshaller) Value() (driver.Value, error) {
	if message.Message == nil {
		return nil, nil
	}

	blob, err := proto.Marshal(message.Message)
	if err != nil {
		return nil, fmt.Errorf(
			"error marshalling protobuf message to blob: %w", err,
		)
	}

	return blob, nil
}

// Scan converts the value from time.Time for the sql driver to handle.
func (message *MessageBlobMarshaller) Scan(src interface{}) error {
	if src == nil {
		*message = MessageBlobMarshaller{}
		return nil
	}

	blob, ok := src.([]byte)
	if !ok {
		return newScanTypeErr(
			blob,
			src,
			message.Message,
		)
	}

	err := proto.Unmarshal(blob, message.Message)
	if err != nil {
		return fmt.Errorf("error unmarshalling message: %w", err)
	}

	return nil
}

// Message creates a new MessageBlobMarshaller for storing / reading an entire protobuf
// message to / from a blob field.
func Message(message proto.Message) *MessageBlobMarshaller {
	return &MessageBlobMarshaller{Message: message}
}
