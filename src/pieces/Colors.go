package pieces

type Color int

const (
	Black Color = 0
	White Color = 1
)

func ToColor(i int) Color {
	var actualIndex = i % 2

	switch actualIndex {
	case 0:
		return Black
	case 1:
		return White
	default:
		return Black
	}
}
