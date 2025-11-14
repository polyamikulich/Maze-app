package domain

type CellType int

const (
	// CellWall представляет стены (#)
	CellWall CellType = iota
	// CellEmpty представляет пустые клетки ( )
	CellEmpty
	// CellStart представляет клетки начала пути (O)
	CellStart
	// CellEnd представляет клетки конца пути (X)
	CellEnd
	// CellPath представляет клетку пути (.)
	CellPath
	// CellSand - песок (~)
	CellSand
	// CellCoin - монетка ($)
	CellCoin
	// CellSmooth - гладкая поверхность (=)
	CellSmooth
	// CellSwamp - болото (&)
	CellSwamp
)

func GetCellWeight(typeCell CellType) int {
	switch typeCell {
	case CellEmpty:
		return 10 // Обычная поверхность
	case CellSand:
		return 15 // Песок - замедляет
	case CellSwamp:
		return 20 // Болото - замедляет
	case CellCoin:
		return 5 // Монетка - ускоряет
	case CellSmooth:
		return 1 // Гладкая поверхность - ускоряет
	case CellPath:
		return 10 // Путь - обычная скорость
	case CellStart:
		return 10 // Старт - обычная скорость
	case CellEnd:
		return 10 // Финиш - обычная скорость
	default:
		return 10 // По умолчанию
	}
}
