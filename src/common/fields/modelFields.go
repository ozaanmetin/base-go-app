package fields

import (
	"database/sql/driver"
	"encoding/json"
)

// Custom type to handle jsonb
type Jsonb map[string]interface{}

func (j *Jsonb) Scan(value interface{}) error {
	// Scan value into Jsonb type
	return json.Unmarshal([]byte(value.([]byte)), j)
}

func (j Jsonb) Value() (driver.Value, error) {
	// Convert Jsonb to JSON string
	return json.Marshal(j)
}
