package common

import "time"

// 做一个另名
type Gender int

const (
	//女性
	Female Gender = 0
	//男性
	Male Gender = 1
	//第三性别
	Third Gender = 2
)

type TimeLayout string

const Data TimeLayout = "2006-01-02"
const DataTime TimeLayout = "2006-01-02 15:04:05"
const DataTimeZone TimeLayout = "2016-01-02 15:04:05Z0700"

func StrToTime(str string, layout ...TimeLayout) (time.Time, error) {
	format := Data
	if len(layout) > 0 {
		format = layout[0]
	}
	return time.Parse(string(format), str)
}
func TimeToStr(t time.Time, layout ...TimeLayout) string {
	format := Data
	if len(layout) > 0 {
		format = layout[0]
	}
	return t.Format(string(format))
}
