package common

import "go.mongodb.org/mongo-driver/bson/bsonrw"

func ConvertToBytesReader(reader interface{}) bsonrw.BytesReader {
	return reader.(bsonrw.BytesReader)
}
