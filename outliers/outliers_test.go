package spe_test

import (
	"reflect"
	"testing"
)

func Test_outliers(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "outliers",
			args: args{slice: []int{2, 4, 0, 100, 4, 11, 2602, 36}},
			want: 11,
		},
		{
			name: "outliers",
			args: args{slice: []int{160, 3, 1719, 19, 11, 13, -21}},
			want: 160,
		},
		{
			name: "outliers",
			args: args{slice: []int{11, 13, 15, 19, 9, 13, -21}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := outliers(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("outliers() = %v, want %v", got, tt.want)
			}
		})
	}
}
