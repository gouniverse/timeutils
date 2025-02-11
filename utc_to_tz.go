package timeutils

import "github.com/dromara/carbon/v2"

// UtcToTZ takes a the provided UTC date and time string and converts to the specified timezone.
//
// Parameters:
// - utcDateTimeString: a string representing the UTC date and time.
// - timezone: a string representing the desired timezone.
//
// Returns:
// - a carbon.Carbon object representing the converted date and time in the specified timezone.
func UtcToTz(utcDateTimeString string, timezone string) carbon.Carbon {
	return carbon.Parse(utcDateTimeString, carbon.UTC).SetTimezone(timezone)
}
