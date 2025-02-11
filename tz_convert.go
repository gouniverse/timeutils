package timeutils

import "github.com/dromara/carbon/v2"

func TzConvert(dateTimeString string, fromTimezone string, toTimezone string) carbon.Carbon {
	return carbon.Parse(dateTimeString, fromTimezone).SetTimezone(toTimezone)
}
