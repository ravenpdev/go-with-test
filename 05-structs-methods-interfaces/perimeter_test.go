package structsmethodsinterfaces

import "testing"

// func TestArea(t *testing.T) {
// 	checkArea := func(t testing.TB, shape Shape, want float64) {
// 		t.Helper()

// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("got %g, want %g", got, want)
// 		}
// 	}

// 	t.Run("rectangles", func(t *testing.T) {
// 		rectangle := Rectangle{10.0, 10.0}
// 		want := 100.0

// 		checkArea(t, rectangle, want)
// 	})

// 	t.Run("cicles", func(t *testing.T) {
// 		circle := Circle{10}
// 		want := 314.1592653589793
// 		checkArea(t, circle, want)
// 	})
// }

// Table Driven Test
func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {

		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt, got, tt.hasArea)
			}
		})

	}
}

// And you can run specific tests within your table with go test -run TestArea/Rectangle.
