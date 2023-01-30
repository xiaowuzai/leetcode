package main

import (
	"reflect"
	"testing"
)

func Test_searchRange1(t *testing.T) {
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
			name:"1",
			args: args{
				nums: []int{5,7,7,8,8,10},
				target:8,
			},
			want: []int{3, 4},
		},
		{
			name:"2",
			args: args{
				nums: []int{5,7,7,8,8,10},
				target:6,
			},
			want: []int{-1, -1},
		},
		{
			name:"1",
			args: args{
				nums: []int{},
				target:0,
			},
			want: []int{-1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange2(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange1() = %v, want %v", got, tt.want)
			}
		})
	}
}
