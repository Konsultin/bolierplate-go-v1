package timek

const (
	// === Date Formats ===

	// FormatDate represents "DD/MM/YY" (e.g., 15/12/24)
	FormatDate = "d/m/y"

	// FormatDateStructure represents "YYYY-MM-DD" (e.g., 2024-12-15)
	// Standard ISO Date format
	FormatDateStructure = "Y-m-d"

	// FormatDateUnStructure represents "DD-MM-YYYY" (e.g., 15-12-2024)
	FormatDateUnStructure = "d-m-Y"

	// FormatDateFull represents "DD MMMM YYYY" (e.g., 15 Desember 2024)
	FormatDateFull = "d F Y"

	// FormatDateMonthYear represents "MMMM YYYY" (e.g., Desember 2024)
	FormatDateMonthYear = "F Y"

	// FormatDateDay represents "Day, DD MMMM YYYY" (e.g., Minggu, 15 Desember 2024)
	FormatDateDay = "l, d F Y"

	// FormatDateShortMonth represents "DD MMM YYYY" (e.g., 15 Des 2024)
	FormatDateShortMonth = "d M Y"

	// === DateTime Formats ===

	// FormatDateTime represents "DD/MM/YY HH:mm" (e.g., 15/12/24 14:30)
	FormatDateTime = "d/m/y H:i"

	// FormatDateTimeStandard represents "YYYY-MM-DD HH:mm:ss" (e.g., 2024-12-15 14:30:00)
	// Standard SQL/Database format
	FormatDateTimeStandard = "Y-m-d H:i:s"

	// FormatDateTimeStructure represents "YYYY-MM-DD HH:mm" (e.g., 2024-12-15 14:30)
	FormatDateTimeStructure = "Y-m-d H:i"

	// FormatDateTimeFull represents "DD MMMM YYYY, HH:mm" (e.g., 15 Desember 2024, 14:30)
	FormatDateTimeFull = "d F Y, H:i"

	// FormatDateTimeDay represents "Day, DD MMMM YYYY HH:mm" (e.g., Minggu, 15 Desember 2024 14:30)
	FormatDateTimeDay = "l, d F Y H:i"

	// FormatDateTimeISO represents "YYYY-MM-DDTHH:mm:ss.000Z"
	// ISO8601 format
	FormatDateTimeISO = "Y-m-d\\TH:i:s.u\\Z"

	// === Time Formats ===

	// FormatTime represents "HH:mm" (e.g., 14:30)
	FormatTime = "H:i"

	// FormatTimeFull represents "HH:mm:ss" (e.g., 14:30:59)
	FormatTimeFull = "H:i:s"

	// === Month Formats ===

	// FormatMonthNumber represents "M" (e.g., 3)
	FormatMonthNumber = "n"

	// FormatMonthNumberZero represents "MM" (e.g., 03)
	FormatMonthNumberZero = "m"

	// FormatMonthShort represents "MMM" (e.g., Mar)
	FormatMonthShort = "M"

	// FormatMonthFull represents "MMMM" (e.g., Maret)
	FormatMonthFull = "F"
)
