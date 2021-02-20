package protosql_test

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"testing"
)

// GetTestDatabase gets a new in-memory sql database, and fails the test if an error
// occurs.
func GetTestDatabase(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if !assert.NoError(t, err, "open database") {
		t.FailNow()
	}

	return db
}

// TestCaseRoundTrip runs a round-trip test-case for
type TestCaseRoundTrip struct {
	// The name of the sub-test.
	Name string
	// Value is the value to encode.
	Value interface{}
	// Decoded is passed to row.Scan.
	Decoded interface{}
	// SQLFieldType is the field type to declare for value.
	SQLFieldType string
	// Expected encoding error
	EncodeErr error
	// Expected decoding error
	DecodeErr error
	// SubTest is an optional additional testing func for inspecting values.
	SubTest func(t *testing.T, testCase *TestCaseRoundTrip)
}

// RunTestRoundTrip runs a round-trip marshalling test based on TestCaseRoundTrip.
func RunTestRoundTrip(t *testing.T, testCase *TestCaseRoundTrip) {
	t.Run(testCase.Name, func(t *testing.T) {
		assert := assert.New(t)

		db := GetTestDatabase(t)

		_, err := db.Exec(
			// QUERY
			fmt.Sprintf(
				"CREATE TABLE test ( value %v );INSERT INTO test (value) VALUES (?)",
				testCase.SQLFieldType,
			),

			// VALUES
			testCase.Value,
		)

		// If we are expecting an error, check that we got it.
		if testCase.EncodeErr != nil {
			assert.EqualError(
				err, testCase.EncodeErr.Error(), "expected encode error",
			)
			// Return, no further tests are needed on expected errors.
			return
			// Otherwise assert that there is no error
		} else if !assert.NoError(err, "create record") {
			t.FailNow()
		}

		row := db.QueryRow(
			`SELECT value FROM test LIMIT 1;`,
		)

		if !assert.NoError(row.Err(), nil) {
			t.FailNow()
		}

		err = row.Scan(testCase.Decoded)
		if testCase.DecodeErr != nil {
			assert.EqualError(
				err, testCase.DecodeErr.Error(), "expected decode error",
			)
			// Return, no further tests are needed on expected errors.
			return
		} else if !assert.NoError(err, "decode value") {
			t.FailNow()
		}

		// Run any additional checking required.
		if testCase.SubTest != nil {
			testCase.SubTest(t, testCase)
		}
	})
}
