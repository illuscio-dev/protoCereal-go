package messagesCereal

import "google.golang.org/protobuf/types/known/timestamppb"

// Mongo db only stores time down to the millisecond, this method clips the accuracy
// of the protobuf time message to match mongo db. This can be useful for testing
// code round trips through BSON / the database.
func ClipTimestamp(timestamp *timestamppb.Timestamp) {
	timestamp.Nanos = timestamp.Nanos / 1e6 * 1e6
}
