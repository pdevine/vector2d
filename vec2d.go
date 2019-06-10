// Package vector2d provides primitives for manipulating vectors on a two dimensional cartesian coordinate system.
package vector2d

import "math"

// A Vec2D is a two dimensional cartesian vector with X and Y values.
type Vec2D struct {
	X, Y float64
}

// NewVec2D creates a new Vec2D vector with given X and Y values.
func NewVec2D(x, y float64) Vec2D {
	return Vec2D{
		X: x,
		Y: y,
	}
}

// IsZero reports whether this is a null vector (i.e. the length of the vector is 0).
func (v Vec2D) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

// IsNorm reports whether this vector is normalized (i.e. the length of the vector is 1).
func (v Vec2D) IsNorm() bool {
	return v.Len() == 1.0
}

// Angle reports the angle of the vector in radians.
func (v Vec2D) Angle() float64 {
	return math.Atan2(v.Y, v.X)
}

// SetAngle creates a new vector at a given angle (given in radians) whose length is the same as the original vector.
func (v Vec2D) SetAngle(a float64) Vec2D {
	l := v.Len()
	v.X = math.Cos(a) * l
	v.Y = math.Sin(a) * l
	return v
}

// SetLen creates a new vector at a given length whose angle is the same as the original vector.
func (v Vec2D) SetLen(l float64) Vec2D {
	a := v.Angle()
	v.X = math.Cos(a) * l
	v.Y = math.Sin(a) * l
	return v
}

// Len reports the length of a vector.
func (v Vec2D) Len() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Norm creates a new vector of length 1 whose angle is the same as the original vector.
func (v Vec2D) Norm() Vec2D {
	l := v.Len()
	if l == 0 {
		return NewVec2D(1, 0)
	}
	v.X /= l
	v.Y /= l
	return v
}

// Trunc creates a new vector a truncated to a given length if the original vector's length is greater than the
// given length.
func (v Vec2D) Trunc(maxl float64) Vec2D {
	l := v.Len()
	return v.SetLen(math.Min(maxl, l))
}

// Reverse creates a new vector in the opposite direction of the original vector.
func (v Vec2D) Reverse() Vec2D {
	v.X = -v.X
	v.Y = -v.Y
	return v
}

// DotProd provides the dot product of two vectors.
func (v Vec2D) DotProd(v2 Vec2D) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

// AngleBetween provides the angle (in radians) between two vectors.
func (v Vec2D) AngleBetween(v2 Vec2D) float64 {
	if !v.IsNorm() {
		v.Norm()
	}
	if !v2.IsNorm() {
		v2.Norm()
	}
	return math.Acos(v.DotProd(v2))
}

// Perp creates a new vector which is perpendicular to the original vector.
func (v Vec2D) Perp() Vec2D {
	v.Y = -v.Y
	return v
}

/*
func (v Vec2D) Sign(v2 Vec2D) int {
	v1 := v.Perp()
	val := v1.DotProd(v2)
	if val < 0 {
		return -1
	}
	return 1
}
*/

// DistSq provides the squared distance between two given vectors.
func (v Vec2D) DistSq(v2 Vec2D) float64 {
	dx := v2.X - v.X
	dy := v2.Y - v.Y
	return dx*dx + dy*dy
}

// Dist provides the distance between two given vectors.
func (v Vec2D) Dist(v2 Vec2D) float64 {
	return math.Sqrt(v.DistSq(v2))
}

// Add creates a new vector by adding two vectors.
func (v Vec2D) Add(v2 Vec2D) Vec2D {
	v.X += v2.X
	v.Y += v2.Y
	return v
}

// Subtract creates a new vector by subtracting two vectors.
func (v Vec2D) Subtract(v2 Vec2D) Vec2D {
	v.X -= v2.X
	v.Y -= v2.Y
	return v
}

// Multiply creates a new vector by multiplying it by a scalar value.
func (v Vec2D) Multiply(val float64) Vec2D {
	v.X *= val
	v.Y *= val
	return v
}

// Divide creates a new vector by dividing it by a scalar value.
func (v Vec2D) Divide(val float64) Vec2D {
	v.X /= val
	v.Y /= val
	return v
}
