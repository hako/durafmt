package durafmt

import (
	"fmt"
	"reflect"
	"strings"
)

// DefaultUnitsCoder default unots coder using `":"` as PluralSep and `","` as UnitsSep
var DefaultUnitsCoder = UnitsCoder{":", ","}

// Unit the pair of singular and plural units
type Unit struct {
	Singular, Plural string
}

// Units durantion units
type Units struct {
	Year, Week, Day, Hour, Minute,
	Second, Millisecond, Microsecond Unit
}

// Units return a slice of units
func (u Units) Units() []Unit {
	return []Unit{u.Year, u.Week, u.Day, u.Hour, u.Minute,
		u.Second, u.Millisecond, u.Microsecond}
}

// UnitsCoder the units encoder and decoder
type UnitsCoder struct {
	// PluralSep char to sep singular and plural pair.
	// Example with char `":"`: `"year:year"` (english) or `"mês:meses"` (portuguese)
	PluralSep,
	// UnitsSep char to sep units (singular and plural pairs).
	// Example with char `","`: `"year:year,week:weeks"` (english) or `"mês:meses,semana:semanas"` (portuguese)
	UnitsSep string
}

// Encode encodes input Units to string
// Examples with `UnitsCoder{PluralSep: ":", UnitsSep = ","}`
// 	- singular and plural pair units: `"year:wers,week:weeks,day:days,hour:hours,minute:minutes,second:seconds,millisecond:millliseconds,microsecond:microsseconds"`
func (coder UnitsCoder) Encode(units Units) string {
	var pairs = make([]string, 8, 8)
	for i, u := range units.Units() {
		pairs[i] = u.Singular + coder.PluralSep + u.Plural
	}
	return strings.Join(pairs, coder.UnitsSep)
}

// Decode decodes input string to Units.
// The input must follow the following formats:
// - Unit format (singular and plural pair)
// 	- must singular (the plural receives 's' character as suffix)
//	- singular and plural: separated by `PluralSep` char
//		Example with char `":"`: `"year:year"` (english) or `"mês:meses"` (portuguese)
// - Units format (pairs of  Year, Week, Day, Hour, Minute,
//	Second, Millisecond and Microsecond units) separated by `UnitsSep` char
// 	- Examples with `UnitsCoder{PluralSep: ":", UnitsSep = ","}`
// 		- must singular units: `"year,week,day,hour,minute,second,millisecond,microsecond"`
// 		- mixed units: `"year,week:weeks,day,hour,minute:minutes,second,millisecond,microsecond"`
// 		- singular and plural pair units: `"year:wers,week:weeks,day:days,hour:hours,minute:minutes,second:seconds,millisecond:millliseconds,microsecond:microsseconds"`
func (coder UnitsCoder) Decode(s string) (units Units, err error) {
	parts := strings.Split(s, coder.UnitsSep)
	if len(parts) != 8 {
		err = fmt.Errorf("bad parts length")
		return
	}
	var ru = reflect.ValueOf(&units).Elem()
	for i, part := range parts {
		ps := strings.Split(part, coder.PluralSep)
		var u Unit
		switch len(ps) {
		case 1:
			// create plural form with sigular + 's' suffix
			u.Singular, u.Plural = ps[0], ps[0]+"s"
		case 2:
			u.Singular, u.Plural = ps[0], ps[1]
		default:
			err = fmt.Errorf("bad unit pair length")
			return
		}
		ru.FieldByIndex([]int{i}).Set(reflect.ValueOf(u))
	}
	return
}
