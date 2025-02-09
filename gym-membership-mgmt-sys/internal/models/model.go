package models

import "time"

type Membership struct {
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    StartDate Date `json:"start_date"`
}

type Date struct {
    time.Time
}

func (d Date) MarshalJSON() ([]byte, error) {
    return []byte(`"` + d.Time.Format("2006-01-02") + `"`), nil
}

// Unmarshal JSON expecting "YYYY-MM-DD"
func (d *Date) UnmarshalJSON(b []byte) error {
    parsed, err := time.Parse(`"2006-01-02"`, string(b))
    if err != nil {
        return err
    }
    d.Time = parsed
    return nil
}