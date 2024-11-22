package customType

import (
	"database/sql/driver"
	"encoding/json"
)

type JsonArray []string

func (t *JsonArray) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

func (t JsonArray) Value() (driver.Value, error) {
	return json.Marshal(t)
}

type JsonMap map[string]interface{}

func (t *JsonMap) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

func (t JsonMap) Value() (driver.Value, error) {
	return json.Marshal(t)
}
