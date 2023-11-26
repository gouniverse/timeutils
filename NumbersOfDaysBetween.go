package timeutils

import "github.com/golang-module/carbon/v2"

func NumberOfDaysBetween(timeStarts string, timeEnds string) int64 {
	numberOfDays := carbon.
		Parse(timeStarts, carbon.UTC).
		DiffInDays(carbon.
			Parse(timeEnds, carbon.UTC))

	return numberOfDays
}
