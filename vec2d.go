package vector2d

import "math"

type Vec2D struct {
	X float64
	Y float64
}

func NewVec2D(x, y float64) Vec2D {
	return Vec2D{
		X: x,
		Y: y,
	}
}

func (v Vec2D) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

func (v Vec2D) IsNorm() bool {
	return v.Len() == 1.0
}

func (v Vec2D) Angle() float64 {
	return math.Atan2(v.Y, v.X)
}

func (v Vec2D) SetAngle(a float64) Vec2D {
	l := v.Len()
	v.X = math.Cos(a) * l
	v.Y = math.Sin(a) * l
	return v
}

func (v Vec2D) SetLen(l float64) Vec2D {
	a := v.Angle()
	v.X = math.Cos(a) * l
	v.Y = math.Sin(a) * l
	return v
}

func (v Vec2D) Len() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vec2D) Norm() Vec2D {
	l := v.Len()
	if l == 0 {
		return NewVec2D(1, 0)
	}
	v.X /= l
	v.Y /= l
	return v
}

func (v Vec2D) Trunc(maxl float64) Vec2D {
	l := v.Len()
	return v.SetLen(math.Min(maxl, l))
}

func (v Vec2D) Reverse() Vec2D {
	v.X = -v.X
	v.Y = -v.Y
	return v
}

func (v Vec2D) DotProd(v2 Vec2D) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

func (v Vec2D) AngleBetween(v2 Vec2D) float64 {
	if !v.IsNorm() {
		v.Norm()
	}
	if !v2.IsNorm() {
		v2.Norm()
	}
	return math.Acos(v.DotProd(v2))
}

// return vec perpendicular to vec
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

func (v Vec2D) DistSq(v2 Vec2D) float64 {
	dx := v2.X - v.X
	dy := v2.Y - v.Y
	return dx*dx + dy*dy
}

func (v Vec2D) Dist(v2 Vec2D) float64 {
	return math.Sqrt(v.DistSq(v2))
}

func (v Vec2D) Add(v2 Vec2D) Vec2D {
	v.X += v2.X
	v.Y += v2.Y
	return v
}

func (v Vec2D) Subtract(v2 Vec2D) Vec2D {
	v.X -= v2.X
	v.Y -= v2.Y
	return v
}

func (v Vec2D) Multiply(val float64) Vec2D {
	v.X *= val
	v.Y *= val
	return v
}

func (v Vec2D) Divide(val float64) Vec2D {
	v.X /= val
	v.Y /= val
	return v
}
