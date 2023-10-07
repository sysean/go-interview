package main

import "testing"

func Test_onReturnAdd(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "匿名返回值",
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := onReturnAdd(); got != tt.want {
				t.Errorf("onReturnAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_onReturnAdd2(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "命名返回值",
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotA := onReturnAdd2(); gotA != tt.want {
				t.Errorf("onReturnAdd2() = %v, want %v", gotA, tt.want)
			}
		})
	}
}
