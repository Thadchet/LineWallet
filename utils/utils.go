package utils

func ToMonthNumber(month string) string {
	switch month {
	case "January":
		return "01"
	case "Fabuary":
		return "02"
	case "March":
		return "03"
	case "April":
		return "04"
	case "May":
		return "05"
	case "June":
		return "06"
	case "July":
		return "07"
	case "August":
		return "08"
	case "Setember":
		return "09"
	case "October":
		return "10"
	case "November":
		return "11"
	case "December":
		return "12"
	default:
		return ""
	}
}

func GetMaxDay(month string) int {
	switch month {
	case "January":
		return 31
	case "Fabuary":
		return 29
	case "March":
		return 31
	case "April":
		return 30
	case "May":
		return 31
	case "June":
		return 30
	case "July":
		return 31
	case "August":
		return 31
	case "Setember":
		return 30
	case "October":
		return 31
	case "November":
		return 30
	case "December":
		return 31
	default:
		return 30
	}
}
