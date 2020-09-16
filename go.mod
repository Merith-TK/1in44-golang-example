module gpio-display

go 1.15

require (
	github.com/stianeikeland/go-rpio/v4 v4.4.0
	lcd v0.0.0-00010101000000-000000000000
)

replace lcd => ./lcd
