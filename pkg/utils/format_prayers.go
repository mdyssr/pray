package utils

import (
	"fmt"
	"github.com/mdyssr/prayers/models"
	"github.com/mdyssr/prayers/names"
	"strings"
)

func formatDate(date models.HijriDate) string {
	formatted := fmt.Sprintf("%s، %s %s %s", date.Weekday.Ar, date.Day, date.Month.Ar, date.Year)
	return formatted
}

func PrintFormattedPrayersData(data models.PrayersData) {
	formattedDate := formatDate(data.HijriDate)
	fmt.Println()
	fmt.Println(formattedDate)
	fmt.Println()

	for _, timing := range data.PrayerTimings {
		name := names.PrayerNames[timing.Name].Ar
		spaces := strings.Repeat(" ", 7-len([]rune(name)))
		designation := timing.Time.Standard.Designation.Ar.Abbreviated
		timingValue := timing.Time.Standard.Value
		formatted := fmt.Sprintf(" %s%s %s %s", name, spaces, timingValue, designation)
		fmt.Println(formatted)
	}

	fmt.Println()

}
