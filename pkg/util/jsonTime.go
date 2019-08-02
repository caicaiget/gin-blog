package util

import (
	"time"
)
type JsonDate time.Time
type JsonTime time.Time

func (p *JsonDate) UnmarshalJSON(data []byte) error {
	local, err := time.ParseInLocation(`"2006-01-02"`, string(data), time.Local)
	*p = JsonDate(local)
	return err
}

func (p *JsonTime) UnmarshalJSON(data []byte) error {
	local, err := time.ParseInLocation(`"2006-01-02 15:04:05"`, string(data), time.Local)
	*p = JsonTime(local)
	return err
}

func (p JsonDate) MarshalJSON() ([]byte, error) {
	data := make([]byte, 0)
	data = append(data, '"')
	data = time.Time(p).AppendFormat(data, "2006-01-02")
	data = append(data, '"')
	return data, nil
}

func (p JsonTime) MarshalJSON() ([]byte, error) {
	data := make([]byte, 0)
	data = append(data, '"')
	data = time.Time(p).AppendFormat(data, "2006-01-02 15:04:05")
	data = append(data, '"')
	return data, nil
}

func (p JsonDate) String() string {
	return time.Time(p).Format("2006-01-02")
}

func (p JsonTime) String() string {
	return time.Time(p).Format("2006-01-02 15:04:05")
}

func Todate(in string) (out time.Time, err error) {
	out, err = time.Parse("2006-01-02", in)
	return out, err
}

func Todatetime(in string) (out time.Time, err error) {
	out, err = time.Parse("2006-01-02 15:04:05", in)
	return out, err
}
