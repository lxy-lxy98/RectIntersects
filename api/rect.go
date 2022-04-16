package api

import (
	"math"
)

type Point2F struct {
	X float64
	Y float64
}

type Rect2F struct {
	Topleft Point2F
	Width   float64
	Height  float64
}

type RectOption interface {
	Intersect() bool //两个矩阵是否相交  包括包含
}

type RectHandler struct {
	DetectRect Rect2F //用于检测的框
	TargetRect Rect2F //目标的框
}

func (rect *RectHandler) Intersect() bool {
	if (math.Abs((rect.TargetRect.Topleft.X+rect.TargetRect.Width/2)-(rect.DetectRect.Topleft.X+rect.DetectRect.Width/2)) <= (rect.DetectRect.Width+rect.TargetRect.Width)/2) && (math.Abs((rect.TargetRect.Topleft.Y-rect.TargetRect.Height/2)-(rect.DetectRect.Topleft.Y-rect.DetectRect.Height/2)) <= (rect.DetectRect.Height+rect.TargetRect.Height)/2) {
		return true
	}
	return false
}
