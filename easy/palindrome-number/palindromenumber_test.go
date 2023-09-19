package easy

import "testing"

func TestIsPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return true when given 121",
			args: args{x: 121},
			want: true,
		},
		{
			name: "should return false when given -121",
			args: args{x: -121},
			want: false,
		},
		{
			name: "should return false when given 10",
			args: args{x: 10},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.x); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
