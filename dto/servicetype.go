package dto

type ServiceType = string

const (
	PPAS_service  ServiceType = "ppas"
	PPAS_notifier ServiceType = "ppas_notifier"
	PPAS_ehed     ServiceType = "ppas_ehed"
	PPAS_updater  ServiceType = "ppas_updater"
)

type SinceTime = string

const (
	WEEK      SinceTime   = "week"
	YESTERDAY SinceTime   = "yesterday"
	HOUR      SinceTime   = "hour"
	MINUTES   ServiceType = "minutes"
)

func IsNotServiceType(input string) bool {
	switch input {
	case PPAS_ehed, PPAS_notifier, PPAS_service, PPAS_updater:
		return true
	}
	return false
}

func IsNotSince(input string) bool {
	switch input {
	case YESTERDAY, HOUR, MINUTES, WEEK:
		return true
	}
	return false
}

func GetSinceArgs(input string) string {
	switch input {
	case MINUTES:
		return "10 minutes ago"
	case HOUR:
		return "2 hours ago"
	case YESTERDAY:
		return "yesterday"
	case WEEK:
		return "1 week ago"
	}
	return "1 week ago"
}
