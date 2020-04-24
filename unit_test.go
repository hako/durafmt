package durafmt

import (
	"reflect"
	"strconv"
	"testing"
)

func TestParseUnits(t *testing.T) {

}

func TestUnitsCoder_Encode(t *testing.T) {
	tests := []struct {
		coder UnitsCoder
		want   string
	}{
		{DefaultUnitsCoder, "year:years,week:weeks,day:days,hour:hours,minute:minutes,second:seconds,millisecond:milliseconds,microsecond:microseconds"},
		{UnitsCoder{"+","|"}, "year+years|week+weeks|day+days|hour+hours|minute+minutes|second+seconds|millisecond+milliseconds|microsecond+microseconds"},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := tt.coder.Encode(units); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitsCoder_Decode(t *testing.T) {
	tests := []struct {
		test    string
		wantU   Units
		wantErr bool
	}{
		{"y,w,d,h,M,s,m,mi", Units{
			Unit{"y", "ys"},
			Unit{"w", "ws"},
			Unit{"d", "ds"},
			Unit{"h", "hs"},
			Unit{"M", "Ms"},
			Unit{"s", "ss"},
			Unit{"m", "ms"},
			Unit{"mi", "mis"},
		}, false},
		{"ano,semana,dia,hora,minuto,segundo,milissegundo,microssegundo", Units{
			Unit{"ano", "anos"},
			Unit{"semana", "semanas"},
			Unit{"dia", "dias"},
			Unit{"hora", "horas"},
			Unit{"minuto", "minutos"},
			Unit{"segundo", "segundos"},
			Unit{"milissegundo", "milissegundos"},
			Unit{"microssegundo", "microssegundos"},
		}, false},
		{"y:YS,w:WS,d:DS,h:HS,M:MS,s:SS,m:mS,mi:MiS", Units{
			Unit{"y", "YS"},
			Unit{"w", "WS"},
			Unit{"d", "DS"},
			Unit{"h", "HS"},
			Unit{"M", "MS"},
			Unit{"s", "SS"},
			Unit{"m", "mS"},
			Unit{"mi", "MiS"},
		}, false},
		{"y:Y:Y_,w,d,h,M,s,m,mi", Units{}, true},
		{"", Units{}, true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			gotU, err := DefaultUnitsCoder.Decode(tt.test)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseUnits() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotU, tt.wantU) {
				t.Errorf("ParseUnits() gotU = %v, want %v", gotU, tt.wantU)
			}
		})
	}
}