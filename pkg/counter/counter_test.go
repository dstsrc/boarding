package counter

import (
	"reflect"
	"testing"
)

func Test_iterate(t *testing.T) {
	type args struct {
		pp []int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 bool
	}{
		{
			name: "first iter",
			args: args{
				[]int{5, 8, 3, 7, 2, 1, 6, 4},
			},
			want:  []int{5, 8, 3, 7, 2, 0, 6, 0},
			want1: false,
		},
		{
			name: "second iter",
			args: args{
				[]int{5, 8, 3, 7, 2, 0, 6, 0},
			},
			want:  []int{5, 8, 3, 7, 0, 0, 0, 0},
			want1: false,
		},
		{
			name: "3 iter",
			args: args{
				[]int{5, 8, 3, 7, 0, 0, 0, 0},
			},
			want:  []int{5, 8, 0, 0, 0, 0, 0, 0},
			want1: false,
		},
		{
			name: "4 iter",
			args: args{
				[]int{5, 8, 0, 0, 0, 0, 0, 0},
			},
			want:  []int{0, 0, 0, 0, 0, 0, 0, 0},
			want1: false,
		},
		{
			name: "5 iter",
			args: args{
				[]int{0, 0, 0, 0, 0, 0, 0, 0},
			},
			want:  []int{0, 0, 0, 0, 0, 0, 0, 0},
			want1: true,
		},
		{
			name: "best case",
			args: args{
				[]int{1, 2, 3, 4, 5},
			},
			want:  []int{0, 0, 0, 0, 0},
			want1: false,
		},
		{
			name: "baddest case 1",
			args: args{
				[]int{3, 2, 1},
			},
			want:  []int{3, 2, 0},
			want1: false,
		},
		{
			name: "baddest case 2",
			args: args{
				[]int{3, 2, 0},
			},
			want:  []int{3, 0, 0},
			want1: false,
		},
		{
			name: "baddest case 3",
			args: args{
				[]int{3, 0, 0},
			},
			want:  []int{0, 0, 0},
			want1: false,
		},
		{
			name: "baddest case 4",
			args: args{
				[]int{0, 0, 0},
			},
			want:  []int{0, 0, 0},
			want1: true,
		},
		{
			name: "test 1",
			args: args{
				[]int{2, 1, 4, 3},
			},
			want:  []int{2, 0, 4, 0},
			want1: false,
		},
		{
			name: "test 2",
			args: args{
				[]int{2, 0, 4, 0},
			},
			want:  []int{0, 0, 0, 0},
			want1: false,
		},
		{
			name: "test 3",
			args: args{
				[]int{3, 4, 1, 2},
			},
			want:  []int{3, 4, 0, 0},
			want1: false,
		},
		{
			name: "test 4",
			args: args{
				[]int{3, 4, 0, 0},
			},
			want:  []int{0, 0, 0, 0},
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := iterate(tt.args.pp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("iterate() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("iterate() got1 = %v, want %v", got1, tt.want1)
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
