package model

import (
    "time"
)

type DateTime time.Time

func (dt *DateTime) UnmarshalJSON(b []byte) error {
    parsedTime, err := time.Parse(`"2006-01-02T15:04:05"`, string(b))
    if err != nil {
        return err
    }
    *dt = DateTime(parsedTime)
    return nil
}

func (dt DateTime) MarshalJSON() ([]byte, error) {
    return []byte(`"` + time.Time(dt).Format("2006-01-02T15:04:05") + `"`), nil
}