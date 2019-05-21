package gollog

import (
	"os"
	"reflect"
	"testing"
)

func TestLog_Start(t *testing.T) {
	type fields struct {
		Level  string
		Thread chan string
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Log{
				Level:  tt.fields.Level,
				Thread: tt.fields.Thread,
			}
			l.Start()
		})
	}
}

func TestOpenLogFile(t *testing.T) {
	type args struct {
		logFile string
	}
	tests := []struct {
		name string
		args args
		want *os.File
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OpenLogFile(tt.args.logFile); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OpenLogFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCloseLogFile(t *testing.T) {
	type args struct {
		log *os.File
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CloseLogFile(tt.args.log)
		})
	}
}

func TestLogIt(t *testing.T) {
	type args struct {
		fyl *os.File
		lyn string
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LogIt(tt.args.fyl, tt.args.lyn)
		})
	}
}

func TestLogOnce(t *testing.T) {
	type args struct {
		logfile string
		msg     string
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LogOnce(tt.args.logfile, tt.args.msg)
		})
	}
}
