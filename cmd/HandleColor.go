package cmd

type COLOR struct {
	name string
	code string
}

func handleColor(C string) string {
	COLORS := []COLOR{
		{"red", ""},
	}
	for _, color := range COLORS {
		if color.name == C {
			return color.code
		}
	}

	return ""
}
