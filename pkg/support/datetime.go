package support

import (
	"fmt"
	"time"
)

const (
	layoutDateTime = "2006-01-02 15:04:05.000"
	layoutTime     = "15:04:05.000"
)

type DateTime struct {
	time.Time
	layout string
}

func ParseDateTime(str string) (DateTime, error) {
	t, err := time.Parse(layoutDateTime, str)
	if err != nil {
		return DateTime{}, err
	}
	return DateTime{Time: t, layout: layoutDateTime}, nil
}

func (d DateTime) String() string {
	return d.Format(d.layout)
}

func ParseTime(str string) (DateTime, error) {
	t, err := time.Parse(layoutTime, str)
	if err != nil {
		return DateTime{}, err
	}
	return DateTime{Time: t, layout: layoutTime}, nil
}

func (d *DateTime) Scan(value any) error {
	var t time.Time
	switch v := value.(type) {
	case time.Time:
		t = v
	case string:
		var err error
		t, err = time.Parse(layoutDateTime, v)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
	*d = DateTime{Time: t, layout: layoutDateTime}
	return nil
}
