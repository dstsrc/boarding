package counterll

import (
	"reflect"
	"testing"
)

func Test_toLL(t *testing.T) {
	type args struct {
		pp []int
	}
	tests := []struct {
		name string
		args args
		want *node
	}{
		{
			name: "simple",
			args: args{
				[]int{1, 2, 3},
			},
			want: &node{value: 3,
				next: &node{value: 2,
					next: &node{value: 1,
						next: nil},
				},
			},
		},
		{
			name: "simple2",
			args: args{
				[]int{3},
			},
			want: &node{value: 3,
				next: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toLL(tt.args.pp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toLL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_iterate(t *testing.T) {
	type args struct {
		head *node
	}
	tests := []struct {
		name string
		args args
		want *node
	}{
		{
			name: "t1",
			args: args{
				head: toLL([]int{5, 8, 3, 7, 2, 1, 6, 4}), // 5, 8, 3, 7, 2, 1, 6, 4
			},
			want: &node{value: 6,
				next: &node{value: 2,
					next: &node{value: 7,
						next: &node{value: 3,
							next: &node{value: 8,
								next: &node{value: 5}}}}}},
		},
		{
			name: "t2",
			args: args{
				head: &node{value: 6,
					next: &node{value: 2,
						next: &node{value: 7,
							next: &node{value: 3,
								next: &node{value: 8,
									next: &node{value: 5}}}}}},
			},
			want: &node{value: 7,
				next: &node{value: 3,
					next: &node{value: 8,
						next: &node{value: 5}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := iterate(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("iterate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_arrToArr(t *testing.T) {
	type args struct {
		pp []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first iter",
			args: args{
				[]int{5, 8, 3, 7, 2, 1, 6, 4},
			},
			want: []int{5, 8, 3, 7, 2, 6},
		},
		{
			name: "second iter",
			args: args{
				[]int{5, 8, 3, 7, 2, 6},
			},
			want: []int{5, 8, 3, 7},
		},
		{
			name: "3 iter",
			args: args{
				[]int{5, 8, 3, 7},
			},
			want: []int{5, 8},
		},
		{
			name: "4 iter",
			args: args{
				[]int{5, 8},
			},
			want: []int{},
		},
		{
			name: "5 iter",
			args: args{
				[]int{},
			},
			want: []int{},
		},
		{
			name: "best case",
			args: args{
				[]int{1, 2, 3, 4, 5},
			},
			want: []int{},
		},
		{
			name: "baddest case 1",
			args: args{
				[]int{3, 2, 1},
			},
			want: []int{3, 2},
		},
		{
			name: "baddest case 2",
			args: args{
				[]int{3, 2},
			},
			want: []int{3},
		},
		{
			name: "baddest case 3",
			args: args{
				[]int{3},
			},
			want: []int{},
		},
		{
			name: "baddest case 4",
			args: args{
				[]int{},
			},
			want: []int{},
		},
		{
			name: "test 1",
			args: args{
				[]int{2, 1, 4, 3},
			},
			want: []int{2, 4},
		},
		{
			name: "test 2",
			args: args{
				[]int{2, 4},
			},
			want: []int{},
		},
		{
			name: "test 3",
			args: args{
				[]int{3, 4, 1, 2},
			},
			want: []int{3, 4},
		},
		{
			name: "test 4",
			args: args{
				[]int{3, 4},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrToArr(tt.args.pp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("arrToArr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCount(t *testing.T) {
	type args struct {
		pp []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "best",
			args: args{
				[]int{1, 2, 3, 4, 5, 6},
			},
			want: 1,
		},
		{
			name: "baddest",
			args: args{
				[]int{6, 5, 4, 3, 2, 1},
			},
			want: 6,
		},
		{
			name: "test 1",
			args: args{
				[]int{5, 8, 3, 7, 2, 1, 6, 4},
			},
			want: 4,
		},
		{
			name: "test 2",
			args: args{
				[]int{9, 8, 7, 6, 5, 1, 2, 3, 4},
			},
			want: 6,
		},
		{
			name: "test 3",
			args: args{
				[]int{8, 1, 6, 4, 2, 5, 3, 7, 9},
			},
			want: 5,
		},
		{
			name: "test 4",
			args: args{
				[]int{2, 3, 4, 1, 7, 8, 9, 5},
			},
			want: 2,
		},
		{
			name: "test 5",
			args: args{
				[]int{2, 3, 4, 1, 9, 8, 7, 5},
			},
			want: 4,
		},
		{
			name: "test 6",
			args: args{
				[]int{1, 5, 2, 3, 4},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Count(tt.args.pp); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}
