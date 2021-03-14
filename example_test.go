package durafmt

import (
	"fmt"
	"math"
	"time"
)

func ExampleParseString() {
	duration, err := ParseString("354h22m3.24s")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(duration) // 2 weeks 18 hours 22 minutes 3 seconds
	// duration.String() // String representation. "2 weeks 18 hours 22 minutes 3 seconds"
}

func ExampleDurafmt_LimitFirstN() {
	duration, err := ParseString("354h22m3.24s")
	if err != nil {
		fmt.Println(err)
	}
	duration = duration.LimitFirstN(2)
	fmt.Println(duration) // 2 weeks 18 hours
	// duration.String() // String representation. "2 weeks 18 hours"
}

func ExampleDurafmt_LimitToUnit() {
	duration, err := ParseString("354h22m3.24s")
	if err != nil {
		fmt.Println(err)
	}
	duration = duration.LimitToUnit("days")
	fmt.Println(duration) // 14 days 18 hours 22 minutes 3 seconds
	// duration.String() // String representation. "14 days 18 hours 22 minutes 3 seconds"
}

func ExampleParseString_sequence() {
	for hours := 1.0; hours < 12.0; hours++ {
		hour := fmt.Sprintf("%fh", math.Pow(2, hours))
		duration, err := ParseString(hour)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(duration) // 2 hours, 4 hours, ...
	}
}

// Version of durafmt.ParseString() that only returns the first part of the duration string.
func ExampleParseStringShort() {
	duration, err := ParseStringShort("354h22m3.24s")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(duration) // 2 weeks 18 hours 22 minutes 3 seconds
	// duration.String() // String representation. "2 weeks 18 hours 22 minutes 3 seconds"
}

func ExampleParse() {
	timeduration := (354 * time.Hour) + (22 * time.Minute) + (3 * time.Second)
	duration := Parse(timeduration).String()
	fmt.Println(duration) // 2 weeks 18 hours 22 minutes 3 seconds
}

// Version of durafmt.Parse() that only returns the first part of the duration string.
func ExampleParseShort() {
	timeduration := (354 * time.Hour) + (22 * time.Minute) + (3 * time.Second)
	duration := ParseShort(timeduration).String()
	fmt.Println(duration) // 2 weeks
}

func ExampleDurafmt_Format() {
	timeduration := (354 * time.Hour) + (22 * time.Minute) + (1 * time.Second) + (100 * time.Microsecond)
	duration := Parse(timeduration)
	// units in portuguese
	units, err := DefaultUnitsCoder.Decode("ano,semana,dia,hora,minuto,segundo,milissegundo,microssegundo")
	if err != nil {
		panic(err)
	}
	fmt.Println(duration.Format(units)) // 2 semanas 18 horas 22 minutos 1 segundo 100 microssegundos

	// custom plural (singular:plural)
	units, err = DefaultUnitsCoder.Decode("ano,semana:SEMANAS,dia,hora,minuto,segundo,milissegundo,microssegundo")
	if err != nil {
		panic(err)
	}
	fmt.Println(duration.Format(units)) // 2 SEMANAS 18 horas 22 minutos 1 segundo 100 microssegundos
}
