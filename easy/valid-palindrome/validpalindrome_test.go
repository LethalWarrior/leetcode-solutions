package validpalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return true when given 'A man, a plan, a canal: Panama'",
			args: args{s: "A man, a plan, a canal: Panama"},
			want: true,
		},
		{
			name: "should return true when given 'race a car'",
			args: args{s: "race a car"},
			want: false,
		},
		{
			name: "should return true when given empty string",
			args: args{s: " "},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.s); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
