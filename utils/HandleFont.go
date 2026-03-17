package utils

func HandleFont(str string) string {
	fonts := []struct {
		name string
		path string
	}{
		{"shadow", "utils/shadow.ascii"},
		{"standard", "utils/standard.ascii"},
		{"thinkertoy", "utils/thinkertoy.ascii"},
	}

	for _, font := range fonts {
		if str == font.name {
			return font.path
		}

	}
	return "utils/standard.ascii"

}
