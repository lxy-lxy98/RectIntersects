package rect_test

import (
	"fmt"
	"testing"

	"github.com/lxy-lxy98/RectIntersects/api"
)

func TestRectContain(t *testing.T) { //判断两个矩阵是否是包含关系
	rect := new(api.RectHandler)
	rect.DetectRect.Topleft.X = 2
	rect.DetectRect.Topleft.Y = 6
	rect.DetectRect.Width = 5
	rect.DetectRect.Height = 4

	// rect.TargetRect.Topleft.X = 4
	// rect.TargetRect.Topleft.Y = 5
	// rect.TargetRect.Width = 1
	// rect.TargetRect.Height = 2

	rect.TargetRect.Topleft.X = 8
	rect.TargetRect.Topleft.Y = 5
	rect.TargetRect.Width = 1
	rect.TargetRect.Height = 2

	if rect.Intersect() {
		fmt.Println("is contain")
	} else {
		fmt.Println("no contain")
	}
}

func TestIntersect(t *testing.T) {
	rect := new(api.RectHandler)
	rect.DetectRect.Topleft.X = 2
	rect.DetectRect.Topleft.Y = 5
	rect.DetectRect.Width = 3
	rect.DetectRect.Height = 3

	rect.TargetRect.Topleft.X = 2
	rect.TargetRect.Topleft.Y = 4
	rect.TargetRect.Width = 2
	rect.TargetRect.Height = 3

	if rect.Intersect() {
		fmt.Println("is intersect")
	} else {
		fmt.Println("no intersect")
	}
}
