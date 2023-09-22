package longestcommonprefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `should return "fl" when given ["flower","flow","flight"]`,
			args: args{strs: []string{"flower", "flow", "flight"}},
			want: "fl",
		},
		{
			name: `should return "" when given ["dog","racecar","car"]`,
			args: args{strs: []string{"dog", "racecar", "car"}},
			want: "",
		},
		{
			name: `should return "" when given [""]`,
			args: args{strs: []string{""}},
			want: "",
		},
		{
			name: `should return "a" when given ["ab","a"]`,
			args: args{strs: []string{"ab", "a"}},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("LongestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
