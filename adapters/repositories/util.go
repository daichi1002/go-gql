package repositories

import "time"

var CustomNow = time.Now

func FakeNow(startDate string, endDate string) {
	CustomNow = func() time.Time {
		fakeTime, _ := time.Parse(startDate, endDate)
		return fakeTime
	}
}
