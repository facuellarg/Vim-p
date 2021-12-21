package entities

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type CustomTime time.Time

var formatDate = `"2006-01-02"`

//MarshalJSON
func (mt CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(mt).Format(formatDate)), nil
}

//UnmarshalJSON
func (mt *CustomTime) UnmarshalJSON(data []byte) error {
	timeFormated, err := time.Parse(formatDate, strings.Trim(string(data), ""))
	if err != nil {
		return err
	}
	*mt = CustomTime(timeFormated)
	return nil
}

func (mt CustomTime) String() string {
	return time.Time(mt).String()
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (mt *CustomTime) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	result := CustomTime{}
	err := json.Unmarshal(bytes, &result)
	*mt = CustomTime(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (mt CustomTime) Value() (driver.Value, error) {
	//if len(j) == 0 {
	//return nil, nil
	//}
	value, err := json.Marshal(mt)
	fmt.Printf("holaa %+v,%+v\n", strings.Trim(string(value), `"`), err)

	return strings.Trim(string(value), `"`), err
}
