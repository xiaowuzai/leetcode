package main

import "testing"

func Test_minSubArrayLen1(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"1",
			args: args{
				target: 7,
				nums:[]int{2,3,1,2,4,3},
			},
			want: 2,
		},
		{
			name:"2",
			args: args{
				target: 4,
				nums:[]int{1,4,4},
			},
			want: 1,
		},
		{
			name:"3",
			args: args{
				target: 11,
				nums:[]int{1,1,1,1,1,1,1,1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen2(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen1() = %v, want %v", got, tt.want)
			}
		})
	}
}
