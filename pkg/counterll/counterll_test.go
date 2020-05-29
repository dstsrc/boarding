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
			name: "first iter",
			args: args{
				[]int{5, 8, 3, 7, 2, 6},
			},
			want: []int{5, 8, 3, 7},
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
