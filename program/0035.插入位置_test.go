package main

import "testing"

func Test_searchInsert1(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{ {
		name: "1",
		args: args{
			nums: []int{1,3,5,6},
			target: 2,
		},
		want: 1,
	}, {
		name: "2",
		args: args{
			nums: []int{1,3,5,6},
			target: 7,
		},
		want: 4,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert1(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_searchInsert2(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{ {
		name: "1",
		args: args{
			nums: []int{1,3,5,6},
			target: 2,
		},
		want: 1,
	}, {
		name: "2",
		args: args{
			nums: []int{1,3,5,6},
			target: 7,
		},
		want: 4,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert2(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_searchInsert3(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{ {
		name: "1",
		args: args{
			nums: []int{1,3,5,6},
			target: 2,
		},
		want: 1,
	}, {
		name: "2",
		args: args{
			nums: []int{1,3,5,6},
			target: 7,
		},
		want: 4,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert3(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
