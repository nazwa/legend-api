package legend

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

// LegendTime represents the most common datetime format used by Legend API
// You'll still find other date formats elsewhere in the API (especially when posting data),
// But this type should cover most GETs
type LegendTime struct {
	time.Time
}

func (b LegendTime) Value() (driver.Value, error) {
	return b.Format("2006-01-02 15:04:05"), nil
}

func (b *LegendTime) Scan(src interface{}) error {
	v, ok := src.(time.Time)

	if !ok {
		return errors.New("Can't scan into LegendTime")
	}

	b.Time = v
	return nil
}

// String "01/02/2006"
func (b LegendTime) String() string {
	return b.Format("01/02/2006")
}

// FormatShortDate "2006-01-02"
func (b LegendTime) FormatShortDate() string {
	return b.Format(LEGEND_DATE_FORMAT)
}

// MarshalJSON "Mon Jan _2"
func (t LegendTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", t.Format("Mon Jan _2"))
	return []byte(stamp), nil
}

// UnmarshalJSON expects either "2006-01-02T15:04:05" or "Mon Jan _2"
func (b *LegendTime) UnmarshalJSON(src []byte) (err error) {
	str := string(src[1 : len(src)-1])

	if str == "ul" {
		return
	}

	b.Time, err = time.Parse("2006-01-02T15:04:05", str)

	if err != nil {
		// Try again with another time format
		b.Time, err = time.Parse("Mon Jan _2", str)
	}
	return
}
