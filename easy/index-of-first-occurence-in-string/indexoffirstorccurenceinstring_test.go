package indexoffirstoccurenceinstring

import "testing"

func TestStrStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `should return 0 when given haystack = "leetcode", needle = "leeto"`,
			args: args{haystack: "leetcode", needle: "leeto"},
			want: -1,
		},
		{
			name: `should return 0 when given haystack = "sadbutsad", needle = "sad"`,
			args: args{haystack: "sadbutsad", needle: "sad"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("StrStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
