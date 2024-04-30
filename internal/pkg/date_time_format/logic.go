package date_time_format

import (
	"strings"
	"time"
)

func GetEndDt(Dt string, Value int) (string, error) {
	StartDt := Dt[0:4] + "-" + Dt[4:6] + "-" + Dt[6:8]
	Dt1, err := time.Parse(DateFormatWithDash2, StartDt)
	if err != nil {
		return "", err
	}
	Dt2 := Dt1.AddDate(Value, 0, 0)
	Dt2 = Dt2.AddDate(0, 0, -1)
	return strings.Replace(Dt2.Format(DateFormatWithDash2), "-", "", -1), nil
}
