package into_test

import (
	"math"
	"testing"

	. "github.com/zenless-lab/into"
)

func TestTryIntoFloat32(t *testing.T) {
	{
		namePrefix := "Float64"
		tests := []struct {
			name    string
			input   float64
			want    float32
			wantErr bool
		}{
			{"positive", 123.456, 123.456, false},
			{"negative", -123.456, -123.456, false},
			{"zero", 0, 0, false},
			{"maxOri", math.MaxFloat64, 0, true},
			{"minOri", -math.MaxFloat64 - 1, 0, true},
			{"maxTarget", float64(math.MaxFloat32), float32(math.MaxFloat32), false},
			{"minTarget", float64(-math.MaxFloat32), float32(-math.MaxFloat32 - 1), false},
			{"maxTarget+1", float64(math.MaxFloat32) * 1.000001, 0, true},  // +1 will not effect the result, because float32 precision
			{"minTarget-1", float64(-math.MaxFloat32) * 1.000001, 0, true}, // -1 will not effect the result, because float32 precision
			{"positiveInf", math.Inf(1), 0, true},
			{"negativeInf", math.Inf(-1), 0, true},
		}

		for _, tt := range tests {
			t.Run(namePrefix+tt.name, func(t *testing.T) {
				got, err := TryIntoFloat32(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("TryIntoFloat32() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("TryIntoFloat32() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	{
		namePrefix := "Float32"
		tests := []struct {
			name    string
			input   float32
			want    float32
			wantErr bool
		}{
			{"positive", 123.456, 123.456, false},
			{"negative", -123.456, -123.456, false},
			{"zero", 0, 0, false},
			{"maxOri", math.MaxFloat32, math.MaxFloat32, false},
			{"minOri", -math.MaxFloat32 - 1, -math.MaxFloat32 - 1, false},
		}

		for _, tt := range tests {
			t.Run(namePrefix+tt.name, func(t *testing.T) {
				got, err := TryIntoFloat32(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("TryIntoFloat32() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("TryIntoFloat32() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	{
		namePrefix := "Bool"
		tests := []struct {
			name    string
			input   bool
			want    float32
			wantErr bool
		}{
			{"true", true, 1, false},
			{"false", false, 0, false},
		}

		for _, tt := range tests {
			t.Run(namePrefix+tt.name, func(t *testing.T) {
				got, err := TryIntoFloat32(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("TryIntoFloat32() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("TryIntoFloat32() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	{
		namePrefix := "Int"
		tests := []struct {
			name    string
			input   int
			want    float32
			wantErr bool
		}{
			{"positive", 123, 123, false},
			{"negative", -123, -123, false},
			{"zero", 0, 0, false},
			{"maxOri", math.MaxInt32, math.MaxInt32, false},
			{"minOri", math.MinInt32, math.MinInt32, false},
			{"maxTarget", math.MaxInt32, math.MaxInt32, false},
			{"minTarget", math.MinInt32, math.MinInt32, false},
		}

		for _, tt := range tests {
			t.Run(namePrefix+tt.name, func(t *testing.T) {
				got, err := TryIntoFloat32(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("TryIntoFloat32() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("TryIntoFloat32() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	{
		namePrefix := "Uint"
		tests := []struct {
			name    string
			input   uint
			want    float32
			wantErr bool
		}{
			{"positive", 123, 123, false},
			{"zero", 0, 0, false},
			{"maxOri", math.MaxUint32, math.MaxUint32, false},
			{"maxTarget", math.MaxUint32, math.MaxUint32, false},
		}

		for _, tt := range tests {
			t.Run(namePrefix+tt.name, func(t *testing.T) {
				got, err := TryIntoFloat32(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("TryIntoFloat32() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("TryIntoFloat32() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	{
		namePrefix := "Int8"
		tests := []struct {
			name    string
			input   int8
			want    float32
			wantErr bool
		}{
			{"positive", 123, 123, false},
			{"negative", -123, -123, false},
			{"zero", 0, 0, false},
			{"maxOri", math.MaxInt8, math.MaxInt8, false},
			{"minOri", math.MinInt8, math.MinInt8, false},
		}

		for _, tt := range tests {
			t.Run(namePrefix+tt.name, func(t *testing.T) {
				got, err := TryIntoFloat32(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("TryIntoFloat32() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("TryIntoFloat32() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	{
		namePrefix := "Int16"
		tests := []struct {
			name    string
			input   int16
			want    float32
			wantErr bool
		}{
			{"positive", 123, 123, false},
			{"negative", -123, -123, false},
			{"zero", 0, 0, false},
			{"maxOri", math.MaxInt16, math.MaxInt16, false},
			{"minOri", math.MinInt16, math.MinInt16, false},
		}

		for _, tt := range tests {
			t.Run(namePrefix+tt.name, func(t *testing.T) {
				got, err := TryIntoFloat32(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("TryIntoFloat32() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("TryIntoFloat32() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	{
		namePrefix := "Int32"
		tests := []struct {
			name    string
			input   int32
			want    float32
			wantErr bool
		}{
			{"positive", 123, 123, false},
			{"negative", -123, -123, false},
			{"zero", 0, 0, false},
			{"maxOri", math.MaxInt32, math.MaxInt32, false},
			{"minOri", math.MinInt32, math.MinInt32, false},
		}

		for _, tt := range tests {
			t.Run(namePrefix+tt.name, func(t *testing.T) {
				got, err := TryIntoFloat32(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("TryIntoFloat32() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("TryIntoFloat32() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	{
		namePrefix := "Int64"
		tests := []struct {
			name    string
			input   int64
			want    float32
			wantErr bool
		}{
			{"positive", 123, 123, false},
			{"negative", -123, -123, false},
			{"zero", 0, 0, false},
			{"maxOri", math.MaxInt64, float32(math.MaxInt64), false},
			{"minOri", math.MinInt64, float32(math.MinInt64), false},
		}

		for _, tt := range tests {
			t.Run(namePrefix+tt.name, func(t *testing.T) {
				got, err := TryIntoFloat32(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("TryIntoFloat32() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("TryIntoFloat32() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	{
		namePrefix := "String"
		tests := []struct {
			name    string
			input   string
			want    float32
			wantErr bool
		}{
			{"positive", "123.456", 123.456, false},
			{"negative", "-123.456", -123.456, false},
			{"zero", "0", 0, false},
			{"invalid", "abc", 0, true},
			{"empty", "", 0, true},
			{"scientific", "1.23456e2", 123.456, false},
			{"positiveInf", "inf", float32(math.Inf(1)), false},
			{"negativeInf", "-inf", float32(math.Inf(-1)), false},
		}

		for _, tt := range tests {
			t.Run(namePrefix+tt.name, func(t *testing.T) {
				got, err := TryIntoFloat32(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("TryIntoFloat32() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("TryIntoFloat32() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	{
		namePrefix := "Uint"
		tests := []struct {
			name    string
			input   uint
			want    float32
			wantErr bool
		}{
			{"positive", 123, 123, false},
			{"zero", 0, 0, false},
			{"maxOri", math.MaxUint32, math.MaxUint32, false},
		}

		for _, tt := range tests {
			t.Run(namePrefix+tt.name, func(t *testing.T) {
				got, err := TryIntoFloat32(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("TryIntoFloat32() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("TryIntoFloat32() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	{
		namePrefix := "Uint8"
		tests := []struct {
			name    string
			input   uint8
			want    float32
			wantErr bool
		}{
			{"positive", 123, 123, false},
			{"zero", 0, 0, false},
			{"maxOri", math.MaxUint8, math.MaxUint8, false},
		}

		for _, tt := range tests {
			t.Run(namePrefix+tt.name, func(t *testing.T) {
				got, err := TryIntoFloat32(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("TryIntoFloat32() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("TryIntoFloat32() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	{
		namePrefix := "Uint16"
		tests := []struct {
			name    string
			input   uint16
			want    float32
			wantErr bool
		}{
			{"positive", 123, 123, false},
			{"zero", 0, 0, false},
			{"maxOri", math.MaxUint16, math.MaxUint16, false},
		}

		for _, tt := range tests {
			t.Run(namePrefix+tt.name, func(t *testing.T) {
				got, err := TryIntoFloat32(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("TryIntoFloat32() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("TryIntoFloat32() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	{
		namePrefix := "Uint32"
		tests := []struct {
			name    string
			input   uint32
			want    float32
			wantErr bool
		}{
			{"positive", 123, 123, false},
			{"zero", 0, 0, false},
			{"maxOri", math.MaxUint32, math.MaxUint32, false},
		}

		for _, tt := range tests {
			t.Run(namePrefix+tt.name, func(t *testing.T) {
				got, err := TryIntoFloat32(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("TryIntoFloat32() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("TryIntoFloat32() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	{
		namePrefix := "Uint64"
		tests := []struct {
			name    string
			input   uint64
			want    float32
			wantErr bool
		}{
			{"positive", 123, 123, false},
			{"zero", 0, 0, false},
			{"maxOri", math.MaxUint64, float32(math.MaxUint64), false},
		}

		for _, tt := range tests {
			t.Run(namePrefix+tt.name, func(t *testing.T) {
				got, err := TryIntoFloat32(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("TryIntoFloat32() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("TryIntoFloat32() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}
