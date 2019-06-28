package vector2d

import (
	"math"
	"testing"
)

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func TestIsZero(t *testing.T) {
	v := NewVec2D(0, 0)
	if v.IsZero() == false {
		t.Error("Expected zero vector")
	}
}

func TestAngle(t *testing.T) {
	v := NewVec2D(-1, 0)
	a := v.Angle()
	if !almostEqual(a, math.Pi) {
		t.Error("Set angle to Pi radians, got ", a)
	}
}

func TestSetAngle(t *testing.T) {
	v := NewVec2D(1, 0)
	v = v.SetAngle(math.Pi / 2)
	if !almostEqual(v.X, 0) && !almostEqual(v.Y, 1.0) {
		t.Errorf("Set angle to Pi/2 radians, got x = %f y = %f", v.X, v.Y)
	}
}

func TestLen(t *testing.T) {
	a, b, c := 3.0, 4.0, 5.0
	v := NewVec2D(a, b)
	l := v.Len()
	if !almostEqual(l, c) {
		t.Error("Expected 3, 4, 5 triangle, got ", l)
	}
}

func TestSetLen(t *testing.T) {
	a, b := 3.0, 4.0
	e, f, g := 6.0, 8.0, 10.0
	v := NewVec2D(a, b)
	v = v.SetLen(g)
	if !almostEqual(v.X, e) || !almostEqual(v.Y, f) {
		t.Errorf("Expected 6, 8, 10 triangle, got x=%f, y=%f", v.X, v.Y)
	}
}

func TestNormZero(t *testing.T) {
	v := NewVec2D(0, 0)
	z := v.Norm()
	l := z.Len()
	if !almostEqual(l, 1.0) {
		t.Error("Expected normalized vector, got len ", l)
	}
	if !almostEqual(z.X, 1.0) || !almostEqual(z.Y, 0.0) {
		t.Errorf("Expected normalized vector (1,0), got (%f,%f)", z.X, z.Y)
	}
}

func TestNorm(t *testing.T) {
	v := NewVec2D(3.0, 4.0)
	if almostEqual(v.Len(), 1.0) {
		t.Error("Expected non-normalized vector")
	}
	v = v.Norm()
	l := v.Len()
	if !almostEqual(l, 1.0) {
		t.Error("Expected normalized vector, got len ", l)
	}
}

func TestIsNorm(t *testing.T) {
	v := NewVec2D(3.0, 4.0)
	if almostEqual(v.Len(), 1.0) {
		t.Error("Expected non-normalized vector")
	}
	v = v.Norm()
	if !v.IsNorm() {
		t.Error("Expected normalized vector")
	}
}

func TestReverse(t *testing.T) {
	a, b := 3.0, 4.0
	v1 := NewVec2D(a, b)
	v2 := v1.Reverse()
	l1 := v1.Len()
	l2 := v2.Len()
	if !almostEqual(l1, l2) {
		t.Errorf("Expected vector lengths to be equal: %f vs %f", l1, l2)
	}
	if !almostEqual(v2.X, -a) && almostEqual(v2.Y, -b) {
		t.Errorf("Expected vector (%f,%f), got (%f,%f)", -a, -b, v2.X, v2.Y)
	}
	v3 := v2.Reverse()
	if !almostEqual(v3.X, a) && almostEqual(v3.Y, b) {
		t.Errorf("Expected vector (%f,%f), got (%f,%f)", a, b, v3.X, v3.Y)
	}
}

func TestTrunc(t *testing.T) {
	v := NewVec2D(3.0, 4.0)
	if almostEqual(v.Len(), 2.0) {
		t.Error("Expected non-normalized vector")
	}
	v = v.Trunc(2.0)
	l := v.Len()
	if !almostEqual(l, 2.0) {
		t.Error("Expected truncated vector length of 2.0, got len ", l)
	}
}

func TestDotProd(t *testing.T) {
	v1 := NewVec2D(3.0, 4.0)
	v2 := NewVec2D(6.0, 8.0)
	val := v1.DotProd(v2)
	if !almostEqual(val, 50.0) {
		t.Error("Expected dot product of 50.0, got ", val)
	}
}

func TestAngleBetweenNorm(t *testing.T) {
	v1 := NewVec2D(1.0, 0)
	v2 := NewVec2D(-1.0, 0)
	a := v1.AngleBetween(v2)
	if !almostEqual(a, math.Pi) {
		t.Error("Expected angle of math.Pi radians, got ", a)
	}
}

func TestAngleBetween(t *testing.T) {
	v1 := NewVec2D(-2.0, 0)
	v2 := NewVec2D(0, -2.0)
	a := v1.AngleBetween(v2)
	if !almostEqual(a, math.Pi/2) {
		t.Error("Expected angle of math.Pi radians, got ", a)
	}
}

func TestPerpVector(t *testing.T) {
	v1 := NewVec2D(2.0, 1.0)
	v2 := v1.Perp()
	if !almostEqual(v2.X, -v1.Y) || !almostEqual(v2.Y, v1.X) {
		t.Error("Expected perp vector")
	}
	val := v1.DotProd(v2)
	if !almostEqual(val, 0) {
		t.Error("Expect zero dot product for a perpendicular vect, got ", val)
	}
	v1 = NewVec2D(-2.0, -1.0)
	v2 = v1.Perp()
	if !almostEqual(v2.X, 1.0) || !almostEqual(v2.Y, -2.0) {
		t.Errorf("Expected perp vector (%f,%f), got (%f,%f)", 1.0, -2.0, v2.X, v2.Y)
	}
}

func TestSign(t *testing.T) {
	v1 := NewVec2D(-1.0, 1.0)
	v2 := NewVec2D(1.0, 0.0)
	val := v1.Sign(v2)
	if val != -1 {
		t.Error("Expected angle to the left (-1), got ", val)
	}
	val = v2.Sign(v1)
	if val != 1 {
		t.Error("Expected angle to the right (1), got ", val)
	}
}

func TestDist(t *testing.T) {
	a, b := 6.0, 8.0
	c, d := 3.0, 4.0
	e := 5.0

	v1 := NewVec2D(a, b)
	v2 := NewVec2D(c, d)
	dist := v2.Dist(v1)
	if !almostEqual(dist, e) {
		t.Errorf("Expected vector dist is %f, got %f", e, dist)
	}
}

func TestAddVec(t *testing.T) {
	a, b, c := 3.0, 4.0, 10.0
	v1 := NewVec2D(a, b)
	v2 := NewVec2D(a, b)
	v3 := v1.Add(v2)
	l := v3.Len()
	if !almostEqual(l, c) {
		t.Errorf("Expected vector len is %f, got %f", c, l)
	}
}

func TestSubVec(t *testing.T) {
	a, b, c := 3.0, 4.0, 5.0
	e, f := 6.0, 8.0
	v1 := NewVec2D(e, f)
	v2 := NewVec2D(a, b)
	v3 := v1.Subtract(v2)
	l := v3.Len()
	if !almostEqual(l, c) {
		t.Errorf("Expected vector len is %f, got %f", c, l)
	}
}

func TestMultiplyVal(t *testing.T) {
	a, b, c := 3.0, 4.0, 10.0
	v1 := NewVec2D(a, b)
	v2 := v1.Multiply(2.0)
	l := v2.Len()
	if !almostEqual(l, c) {
		t.Errorf("Expected vector len is %f, got %f", c, l)
	}
}

func TestDivideVec(t *testing.T) {
	a, b, c := 6.0, 8.0, 5.0
	v1 := NewVec2D(a, b)
	v2 := v1.Divide(2.0)
	l := v2.Len()
	if !almostEqual(l, c) {
		t.Errorf("Expected vector len is %f, got %f", c, l)
	}
}
