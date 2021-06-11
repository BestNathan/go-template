package time

import (
	"strconv"
	"strings"
	"time"
)

type Timestamp time.Time

func New() Timestamp {
	return Timestamp(time.Now())
}

func (t *Timestamp) UnmarshalJSON(s []byte) (err error) {
	r := strings.Replace(string(s), `"`, ``, -1)
	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.Unix(0, q*1e6)

	return
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).UnixNano()/1e6, 10)), nil
}

func FromTime(t time.Time) Timestamp {
	return Timestamp(t)
}
