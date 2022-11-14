package needle

import "testing"

func Test_needle(t *testing.T) {
	type args struct {
		slice []string
		text  string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "needle",
			args: args{slice: []string{"red", "blue", "yellow", "black", "grey"}, text: "blue"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := needle(tt.args.slice, tt.args.text); got != tt.want {
				t.Errorf("needle() = %v, want %v", got, tt.want)
			}
		})
	}
}
