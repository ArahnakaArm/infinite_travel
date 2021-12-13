package services

import (
	"regexp"
)

func DateTimeValidate(datetime string) bool {

	var digitCheckNumber = regexp.MustCompile(`^[0-9]+$`)

	var digitCheckDash = regexp.MustCompile(`-`)

	var digitCheckSpace = regexp.MustCompile(` `)

	var digitCheckColon = regexp.MustCompile(`:`)

	var digitCheckPoint = regexp.MustCompile(`\.`)

	if len(datetime) != 23 || !digitCheckNumber.MatchString(datetime[0:4]) || !digitCheckDash.MatchString(datetime[4:5]) || !digitCheckNumber.MatchString(datetime[5:7]) ||
		!digitCheckDash.MatchString(datetime[7:8]) || !digitCheckNumber.MatchString(datetime[8:10]) || !digitCheckSpace.MatchString(datetime[10:11]) ||
		!digitCheckNumber.MatchString(datetime[11:13]) || !digitCheckColon.MatchString(datetime[13:14]) || !digitCheckNumber.MatchString(datetime[14:16]) || !digitCheckColon.MatchString(datetime[16:17]) ||
		!digitCheckNumber.MatchString(datetime[17:19]) || !digitCheckPoint.MatchString(datetime[19:20]) || !digitCheckNumber.MatchString(datetime[20:23]) {
		return false
	} else {
		return true
	}

}
