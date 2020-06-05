package utils

import "strconv"

func (u *Utils) Atoi(s string, defaultValue int) int {
	sInt, sIntErr := strconv.Atoi(s)
	if sIntErr != nil {
		if u.debug {
			u.l.Warnf("Unable to convert string \"%s\" to integer. %s", s, sIntErr.Error())
			u.l.Warnf("Using default value %d", defaultValue)
		}
		sInt = defaultValue
	}

	return sInt
}

func (u *Utils) ParseBool(s string, defaultValue bool) bool {
	sBool, sBoolErr := strconv.ParseBool(s)
	if sBoolErr != nil {
		if u.debug {
			u.l.Warnf("Unable to convert string \"%s\" to integer. %s", s, sIntErr.Error())
			u.l.Warnf("Using default value %d", defaultValue)
		}
	}
	return sBool
}