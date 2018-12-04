package timespan

import (
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	type args struct {
		start time.Time
		end   time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    TimeSpan
		wantErr bool
	}{
		{
			name: "start = end",
			args: args{
				start: time.Date(2018, time.December, 5, 0, 0, 0, 0, time.Local),
				end:   time.Date(2018, time.December, 5, 0, 0, 0, 0, time.Local),
			},
			want: TimeSpan{
				start: time.Date(2018, time.December, 5, 0, 0, 0, 0, time.Local),
				end:   time.Date(2018, time.December, 5, 0, 0, 0, 0, time.Local),
			},
		},
		{
			name: "start < end",
			args: args{
				start: time.Date(2018, time.December, 5, 0, 0, 0, 0, time.Local),
				end:   time.Date(2018, time.December, 5, 0, 0, 0, 1, time.Local),
			},
			want: TimeSpan{
				start: time.Date(2018, time.December, 5, 0, 0, 0, 0, time.Local),
				end:   time.Date(2018, time.December, 5, 0, 0, 0, 1, time.Local),
			},
		},
		{
			name: "start > end",
			args: args{
				start: time.Date(2018, time.December, 5, 0, 0, 0, 0, time.Local),
				end:   time.Date(2018, time.December, 4, 0, 0, 0, 0, time.Local),
			},
			want:    TimeSpan{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
