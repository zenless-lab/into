package into

import (
	"testing"
	"time"
)

func TestToInt32ToFloat32(t *testing.T) {
	var value int32 = 12345
	var result float32
	result = Into[float32](value)

	if result != float32(value) {
		t.Fatalf("expected %v, got %v", float32(value), result)
	}
}

func TestToInt32ToFloat32EdgeCase(t *testing.T) {
	var value int32 = 2147483647
	var result float32
	result = Into[float32](value)

	if result != float32(value) {
		t.Fatalf("expected %v, got %v", float32(value), result)
	}
}

func TestToInt32ToFloat32Negative(t *testing.T) {
	var value int32 = -12345
	var result float32
	result = Into[float32](value)

	if result != float32(value) {
		t.Fatalf("expected %v, got %v", float32(value), result)
	}
}

func TestToInt32ToFloat32NegativeEdgeCase(t *testing.T) {
	var value int32 = -2147483648
	var result float32
	result = Into[float32](value)

	if result != float32(value) {
		t.Fatalf("expected %v, got %v", float32(value), result)
	}
}

func TestToInt32ToUnsupportedType(t *testing.T) {
	var value int32 = 12345
	// must panic
	didPanic := false

	defer func() {
		if r := recover(); r != nil {
			didPanic = true
		}

		if !didPanic {
			t.Fatalf("expected panic, got %v", Into[time.Time](value))
		}
	}()

	Into[time.Time](value)
}
