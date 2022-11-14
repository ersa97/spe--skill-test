package blueocean

import (
	"reflect"
	"testing"
)

func Test_blueocean(t *testing.T) {
	type args struct {
		slice1 []int
		slice2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "blueocean",
			args: args{slice1: []int{1, 2, 3, 4, 6, 10}, slice2: []int{1}},
			want: []int{2, 3, 4, 6, 10},
		},
		{
			name: "blueocean",
			args: args{slice1: []int{1, 5, 5, 5, 5, 3}, slice2: []int{5}},
			want: []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := blueocean(tt.args.slice1, tt.args.slice2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("blueocean() = %v, want %v", got, tt.want)
			}
		})
	}
}
