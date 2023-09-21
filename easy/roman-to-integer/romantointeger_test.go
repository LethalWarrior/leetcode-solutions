package romantointeger

import "testing"

func TestRomanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return 3 when given III",
			args: args{s: "III"},
			want: 3,
		},
		{
			name: "should return 58 when given LVIII",
			args: args{s: "LVIII"},
			want: 58,
		},
		{
			name: "should return 1994 when given MCMXCIV",
			args: args{s: "MCMXCIV"},
			want: 1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RomanToInt(tt.args.s); got != tt.want {
				t.Errorf("RomanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
