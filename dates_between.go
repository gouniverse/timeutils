package timeutils

import "github.com/dromara/carbon/v2"

func DatesBetween(timeStarts string, timeEnds string) []string {
	numberOfDays := NumberOfDaysBetween(timeStarts, timeEnds)

	dates := []string{}

	timeStartsCarbon := carbon.Parse(timeStarts, carbon.UTC)

	for i := 0; i <= int(numberOfDays); i++ {
		day := timeStartsCarbon.AddDays(i).Format("Y-m-d")
		dates = append(dates, day)
	}

	return dates
}
