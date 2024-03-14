package main

import "testing"

func Test_get_pow_of_two(t *testing.T) {
	type args struct {
		capacity uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "test1",
			args: args{capacity: 1022},
			want: 1024,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPowOfTwo(tt.args.capacity); got != tt.want {
				t.Errorf("GetPowOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
