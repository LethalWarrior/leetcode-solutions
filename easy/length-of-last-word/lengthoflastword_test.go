package lengthoflastword

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `should return 5 when given "Hello World"`,
			args: args{s: "Hello World"},
			want: 5,
		},
		{
			name: `should return 6 when given "luffy is still joyboy"`,
			args: args{s: "luffy is still joyboy"},
			want: 6,
		},
		{
			name: `should return 4 when given "   fly me   to   the moon  "`,
			args: args{s: "   fly me   to   the moon  "},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("LengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
