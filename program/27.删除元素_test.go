package main

import "testing"

func Test_removeElement1(t *testing.T) {
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
			name: "1",
			args: args{
				nums: []int{3,2,2,3},
				val: 3,
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				nums: []int{0,1,2,2,3,0,4,2},
				val: 2,
			},
			want: 5,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement3(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement1() = %v, want %v", got, tt.want)
			}
		})
	}
}
