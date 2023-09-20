package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "should return [0,1] when given [2,7,11,15] and 9",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "should return [1,2] when given [3,2,4] and 6",
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
		},
		{
			name: "should return [0,1] when given [3,3] and 6",
			args: args{
				nums:   []int{3, 3},
				target: 6,
			},
			want: []int{0, 1},
		},
		{
			name: "should return [0,1] when given [1,9,13,20,47] and 10",
			args: args{
				nums:   []int{1, 9, 13, 20, 47},
				target: 10,
			},
			want: []int{0, 1},
		},
		{
			name: "should return [0,4] when given [3,2,4,1,9] and 12",
			args: args{
				nums:   []int{3, 2, 4, 1, 9},
				target: 12,
			},
			want: []int{0, 4},
		},
		{
			name: "should return [] when given [] and 10",
			args: args{
				nums:   []int{},
				target: 10,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
