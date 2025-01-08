package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalc(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name:    "Simple addition",
			args:    args{expression: "1 + 2"},
			want:    3.0,
			wantErr: false,
		},
		{
			name:    "Simple subtraction",
			args:    args{expression: "5 - 3"},
			want:    2.0,
			wantErr: false,
		},
		{
			name:    "Multiplication and division",
			args:    args{expression: "6 * 2 / 3"},
			want:    4.0,
			wantErr: false,
		},
		{
			name:    "Division by zero",
			args:    args{expression: "4 / 0"},
			want:    0.0,
			wantErr: true,
		},
		{
			name:    "Invalid expression",
			args:    args{expression: "2 +"},
			want:    0.0,
			wantErr: true,
		},
		{
			name:    "Complex expression",
			args:    args{expression: "(1 + 2) * (3 + 4)"},
			want:    21.0,
			wantErr: false,
		},
		{
			name:    "Float numbers",
			args:    args{expression: "3.5 + 4.2"},
			want:    7.7,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calc(tt.args.expression)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
