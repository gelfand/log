package log

import (
	"bytes"
	"fmt"
	"regexp"
	"testing"
)

var reg = regexp.MustCompile("[0-9][0-9]:[0-9][0-9]:[0-9][0-9].[0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9] [a-zA-Z][a-zA-Z]")

func TestLog(t *testing.T) {
	tests := []struct {
		name string
		lvl  logLvl
		str  string
	}{
		{
			name: "info",
			lvl:  lvlInfo,
			str:  "Hello",
		},
		{
			name: "warn",
			lvl:  lvlWarn,
			str:  "Hello",
		},
		{
			name: "error",
			lvl:  lvlError,
			str:  "Hello",
		},
		{
			name: "debug",
			lvl:  lvlDebug,
			str:  "Hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			log := New(buf)
			switch tt.lvl {
			case lvlInfo:
				log.Info(tt.str)
			case lvlWarn:
				log.Warn(tt.str)
			case lvlError:
				log.Error(tt.str)
			case lvlDebug:
				log.Debug(tt.str)
			}
			got := reg.ReplaceAllString(buf.String(), "")
			want := tt.lvl.color() + tt.lvl.String() + tt.lvl.reset() + "[] " + tt.str + "\n"
			if got != want {
				t.Errorf("%s: got %q, want %q", tt.name, got, want)
			}
		})
	}
}

func TestLogf(t *testing.T) {
	tests := []struct {
		name   string
		lvl    logLvl
		str    string
		format []interface{}
	}{
		{
			name:   "infof",
			lvl:    lvlInfo,
			str:    "Hello %d %d",
			format: []interface{}{1, 2},
		},
		{
			name:   "warnf",
			lvl:    lvlWarn,
			str:    "Hello %d %d",
			format: []interface{}{1, 2},
		},
		{
			name:   "errorf",
			lvl:    lvlError,
			str:    "Hello %d %d",
			format: []interface{}{1, 2},
		},
		{
			name:   "debugf",
			lvl:    lvlDebug,
			str:    "Hello %d %d",
			format: []interface{}{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			log := New(buf)
			switch tt.lvl {
			case lvlInfo:
				log.Infof(tt.str, tt.format...)
			case lvlWarn:
				log.Warnf(tt.str, tt.format...)
			case lvlError:
				log.Errorf(tt.str, tt.format...)
			case lvlDebug:
				log.Debugf(tt.str, tt.format...)
			}

			got := reg.ReplaceAllString(buf.String(), "")
			want := tt.lvl.color() + tt.lvl.String() + tt.lvl.reset() + "[] " + fmt.Sprintf(tt.str, tt.format...) + "\n"
			if got != want {
				t.Errorf("%s: got %q, want %q", tt.name, got, want)
			}
		})
	}
}
