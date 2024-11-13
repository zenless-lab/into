package into

import (
	"testing"
)

func TestToInt32ToFloat32(t *testing.T) {
	var value int32 = 12345
	var result float32 = Into[float32](value)

	if result != float32(value) {
		t.Fatalf("expected %v, got %v", float32(value), result)
	}
}

func TestToInt32ToFloat32EdgeCase(t *testing.T) {
	var value int32 = 2147483647
	var result float32 = Into[float32](value)

	if result != float32(value) {
		t.Fatalf("expected %v, got %v", float32(value), result)
	}
}

func TestToInt32ToFloat32Negative(t *testing.T) {
	var value int32 = -12345
	var result float32 = Into[float32](value)

	if result != float32(value) {
		t.Fatalf("expected %v, got %v", float32(value), result)
	}
}

func TestToInt32ToFloat32NegativeEdgeCase(t *testing.T) {
	var value int32 = -2147483648
	var result float32 = Into[float32](value)

	if result != float32(value) {
		t.Fatalf("expected %v, got %v", float32(value), result)
	}
}
