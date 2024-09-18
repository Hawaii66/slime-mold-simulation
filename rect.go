package main

type Rect struct {
	originX int
	originY int
	endX    int
	endY    int
}

func (rect *Rect) IsInside(x, y int) bool {
	if x < rect.originX {
		return false
	}
	if y < rect.originY {
		return false
	}
	if x > rect.endX-1 {
		return false
	}
	if y > rect.endY-1 {
		return false
	}

	return true
}

func (rect *Rect) ClampInside(x, y int) (int, int) {
	newX := x
	newY := y
	if x < rect.originX {
		newX = rect.originX
	}
	if y < rect.originY {
		newY = rect.originY
	}
	if x > rect.endX {
		newX = rect.endX - 1
	}
	if y > rect.endY {
		newY = rect.endY - 1
	}

	return newX, newY
}
