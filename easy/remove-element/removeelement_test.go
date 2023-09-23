package removeelement

import "testing"

func TestRemoveElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return 2 when given nums = [3,2,2,3], val = 3",
			args: args{nums: []int{3, 2, 2, 3}, val: 3},
			want: 2,
		},
		{
			name: "should return 5 when given nums = [0,1,2,2,3,0,4,2], val = 2",
			args: args{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("RemoveElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
