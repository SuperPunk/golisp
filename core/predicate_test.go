package core

import "testing"

func TestTrue(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"t1", args{1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := True(tt.args.v); got != tt.want {
				t.Errorf("True() = %v, want %v", got, tt.want)
			}
		})
	}
}
