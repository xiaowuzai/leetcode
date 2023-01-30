package main

import "testing"

func Test_search1(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"1",
			args: args{
				nums:[]int{-1,0,3,5,9,12},
				target: 9,
			},
			want: 4,
		},
		{
			name:"2",
			args: args{
				nums:[]int{-1,0,3,5,9,12},
				target: 2,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search1(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search1() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_search2(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"1",
			args: args{
				nums:[]int{-1,0,3,5,9,12},
				target: 9,
			},
			want: 4,
		},
		{
			name:"2",
			args: args{
				nums:[]int{-1,0,3,5,9,12},
				target: 2,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search2(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSearch1(b *testing.B) {
	nums := []int{-1,0,3,5,9,12}
	target := 9
	for i := 0; i < b.N; i++ {
		search1(nums, target)
	}
}

func BenchmarkSearch2(b *testing.B) {
	nums := []int{-1,0,3,5,9,12}
	target := 9
	for i := 0; i < b.N; i++ {
		search2(nums, target)
	}
}
