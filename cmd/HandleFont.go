package cmd

type FONT struct {
	name string
	path string
}

func HandleFont(BANNER string) string {
	FONTS := []FONT{
		{"shadow", "banners/shadow.txt"},
		{"standard", "banners/standard.txt"},
		{"thinkertoy", "banners/thinkertoy.txt"},
	}

	for _, font := range FONTS {
		if BANNER == font.name {
			return font.path
		}
	}
	return ""
}
