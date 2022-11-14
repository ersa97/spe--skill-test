package narsis

import "testing"

func Test_narsis(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "narsis",
			args: args{number: 153},
			want: true,
		},
		{
			name: "narsis",
			args: args{number: 111},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := narsis(tt.args.number); got != tt.want {
				t.Errorf("narsis() = %v, want %v", got, tt.want)
			}
		})
	}
}
