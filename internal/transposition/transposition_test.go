package transposition

import (
	"reflect"
	"testing"
)

func Test_getInitParams(t *testing.T) {
	type args struct {
		pp []int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 int
	}{
		{
			name: "transposition",
			args: args{
				pp: []int{6, 5, 3, 1, 2, 4, 8, 7, 9},
			},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want1: 362880,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getInitParams(tt.args.pp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getInitParams() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getInitParams() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_revers(t *testing.T) {
	type args struct {
		ss []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1-9",
			args: args{ss: []int{1, 2, 3, 4, 5, 6, 7, 8}},
			want: []int{8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			name: "empty",
			args: args{[]int{}},
			want: []int{},
		},
		{
			name: "trio",
			args: args{[]int{1, 2, 3}},
			want: []int{3, 2, 1},
		},
		{
			name: "pair",
			args: args{[]int{1, 2}},
			want: []int{2, 1},
		},
		{
			name: "single",
			args: args{[]int{1}},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := revers(tt.args.ss); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("revers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_next1(t *testing.T) {
	ss := []int{1, 2, 3}
	for i := 0; i < 5; i++ {
		next(ss)
	}
	if !reflect.DeepEqual(ss, []int{3, 2, 1}) {
		t.Fatal()
	}

	ss = []int{1, 2, 3, 4}
	for i := 0; i < 23; i++ {
		next(ss)
	}
	if !reflect.DeepEqual(ss, []int{4, 3, 2, 1}) {
		t.Fatal()
	}

	ss = []int{1, 2, 2, 3}
	for i := 0; i < 11; i++ {
		next(ss)
	}
	if !reflect.DeepEqual(ss, []int{3, 2, 2, 1}) {
		t.Fatal()
	}
}

func Test_next(t *testing.T) {
	type args struct {
		ss []int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 bool
	}{

		{
			name:  "base",
			args:  args{[]int{1, 2, 3, 4, 5, 9, 8, 7, 6}},
			want:  []int{1, 2, 3, 4, 6, 5, 7, 8, 9},
			want1: false,
		},
		{
			name:  "3-1",
			args:  args{[]int{1, 2, 3}},
			want:  []int{1, 3, 2},
			want1: false,
		},
		{
			name:  "3-2",
			args:  args{[]int{1, 3, 2}},
			want:  []int{2, 1, 3},
			want1: false,
		},
		{
			name:  "3-3",
			args:  args{[]int{2, 1, 3}},
			want:  []int{2, 3, 1},
			want1: false,
		},
		{
			name:  "3-4",
			args:  args{[]int{2, 3, 1}},
			want:  []int{3, 1, 2},
			want1: false,
		},
		{
			name:  "3-5",
			args:  args{[]int{3, 1, 2}},
			want:  []int{3, 2, 1},
			want1: false,
		},
		{
			name:  "3-6",
			args:  args{[]int{3, 2, 1}},
			want:  nil,
			want1: true,
		},
		{
			name:  "single",
			args:  args{[]int{1}},
			want:  nil,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := next(tt.args.ss)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("next() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("next() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
