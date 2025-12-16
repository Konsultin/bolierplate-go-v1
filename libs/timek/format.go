package timek

import "strings"

// Format returns formatted string using PHP-style format characters (Carbon compatible)
// d: Day of month, 2 digits with leading zeros (01 to 31)
// m: Numeric representation of a month, with leading zeros (01 to 12)
// y: A two digit representation of a year (e.g. 99 or 03)
// Y: A full numeric representation of a year, 4 digits (e.g. 1999 or 2003)
// F: A full textual representation of a month (January through December)
// l: A full textual representation of the day of the week (Sunday through Saturday)
// H: 24-hour format of an hour with leading zeros (00 through 23)
// i: Minutes with leading zeros (00 to 59)
// s: Seconds with leading zeros (00 to 59)
func (t Time) Format(layout string) string {
	return t.Carbon.Format(layout)
}

// ToReadable returns a human-readable relative time string
// e.g., "1 hour ago", "5 menit yang lalu"
// Depends on the locale set via SetLocale()
func (t Time) ToReadable() string {
	return t.DiffForHumans()
}

// ToDate returns "DD/MM/YY" e.g., 15/12/24
func (t Time) ToDate() string {
	return t.Format(FormatDate)
}

// ToDateFull returns "DD MMMM YYYY" e.g., 15 Desember 2024
func (t Time) ToDateFull() string {
	return t.Format(FormatDateFull)
}

// ToDateTime returns "DD/MM/YY HH:mm" e.g., 15/12/24 14:30
func (t Time) ToDateTime() string {
	return t.Format(FormatDateTime)
}

// ToDateDay returns "Day, DD MMMM YYYY" e.g., Minggu, 15 Desember 2024
func (t Time) ToDateDay() string {
	return t.Format(FormatDateDay)
}

// ToSQL returns "YYYY-MM-DD HH:mm:ss"
func (t Time) ToSQL() string {
	return t.Format(FormatDateTimeStandard)
}

// ToTime returns "HH:mm" e.g., 14:30
func (t Time) ToTimeStr() string {
	return t.Format(FormatTime)
}

// ToMonthNumber returns month number without leading zero e.g., "3"
func (t Time) ToMonthNumber() string {
	return t.Format(FormatMonthNumber)
}

// ToMonthNumberZero returns month number with leading zero e.g., "03"
func (t Time) ToMonthNumberZero() string {
	return t.Format(FormatMonthNumberZero)
}

// ToMonthShort returns short month name e.g., "Mar" (localized)
func (t Time) ToMonthShort() string {
	return t.Format(FormatMonthShort)
}

// ToMonthFull returns full month name e.g., "Maret" (localized)
func (t Time) ToMonthFull() string {
	return t.Format(FormatMonthFull)
}

// NormalizeMonthName converts English month names to Indonesian in a string if needed
// Useful for fallback or manual processing
func NormalizeMonthName(s string) string {
	replacer := strings.NewReplacer(
		"January", "Januari",
		"February", "Februari",
		"March", "Maret",
		"May", "Mei",
		"June", "Juni",
		"July", "Juli",
		"August", "Agustus",
		"October", "Oktober",
		"December", "Desember",
	)
	return replacer.Replace(s)
}
