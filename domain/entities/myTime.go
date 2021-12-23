package entities

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

type CustomTime time.Time

var formatDateSql = "2006-01-02"
var formatDate = `"2006-01-02"`

//MarshalJSON
func (mt CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(mt).Format(formatDate)), nil
}

//UnmarshalJSON
func (mt *CustomTime) UnmarshalJSON(data []byte) error {
	return parseMyTime(formatDate, data, mt)
}

func parseMyTime(format string, data []byte, mt *CustomTime) error {
	timeFormated, err := time.Parse(format, strings.Trim(string(data), ""))
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
	//result := CustomTime{}
	err := parseMyTime(formatDateSql, bytes, mt)
	//*mt = CustomTime(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (mt CustomTime) Value() (driver.Value, error) {
	//if len(j) == 0 {
	//return nil, nil
	//}
	value := time.Time(mt).Format(formatDateSql)
	//value, err := json.Marshal(mt)

	return value, nil
}
