package model

import (
	"database/sql/driver"
	"strconv"
	"time"
)

type Time time.Time

func (t Time) String() string {
	return strconv.FormatInt(time.Time(t).Unix(), 10)
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(t.String()), nil
}
func (t Time) Value() (driver.Value, error) {
	return time.Time(t), nil
}

type Times struct {
	CreatedAt Time `json:"created_at"`
	UpdatedAt Time `json:"updated_at"`
}
