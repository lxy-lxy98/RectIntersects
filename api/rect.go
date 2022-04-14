package api

import "fmt"

type Point2F struct {
	x float64
	y float64
}

type Rect2F struct {
	topleft Point2F
	width   float64
	height  float64
}

type RectOption interface {
	Contain() bool
	Intersect() bool
}

type RectHandler struct {
	srcRect Rect2F
	dstRect Rect2F
}

func (rect *RectHandler) Contain() bool {

	return false
}

func (rect *RectHandler) Intersect() bool {
	fmt.Println("lxy hello")
	return false
}
