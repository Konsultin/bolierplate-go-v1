package timek

import (
	"time"

	"github.com/dromara/carbon/v2"
)

const (
	// LangID is Indonesian language code
	LangID = "id"
	// LangEN is English language code
	LangEN = "en"
)

// Time wraps carbon.Carbon to provide enhanced functionality
type Time struct {
	carbon.Carbon
}

// Now returns current time
func Now() Time {
	return Time{
		Carbon: *carbon.Now().SetLanguage(carbon.NewLanguage().SetLocale(LangID)),
	}
}

// Parse parses a time string
func Parse(timeStr string) Time {
	return Time{
		Carbon: *carbon.Parse(timeStr).SetLanguage(carbon.NewLanguage().SetLocale(LangID)),
	}
}

// FromTime creates Time from time.Time
func FromTime(t time.Time) Time {
	return Time{
		Carbon: *carbon.CreateFromStdTime(t).SetLanguage(carbon.NewLanguage().SetLocale(LangID)),
	}
}

// FromUnix creates Time from unix timestamp
func FromUnix(seconds int64) Time {
	return Time{
		Carbon: *carbon.CreateFromTimestamp(seconds).SetLanguage(carbon.NewLanguage().SetLocale(LangID)),
	}
}

// Date creates Time from year, month, day
func Date(year, month, day int) Time {
	return Time{
		Carbon: *carbon.CreateFromDate(year, month, day).SetLanguage(carbon.NewLanguage().SetLocale(LangID)),
	}
}

// SetLocale sets the language for the time instance (affects formatting and readable output)
func (t Time) SetLocale(lang string) Time {
	return Time{
		Carbon: *t.Carbon.SetLanguage(carbon.NewLanguage().SetLocale(lang)),
	}
}

// ToTime returns standard time.Time
func (t Time) ToTime() time.Time {
	return t.Carbon.StdTime()
}

// String returns string representation
func (t Time) String() string {
	return t.ToDateTimeString()
}
