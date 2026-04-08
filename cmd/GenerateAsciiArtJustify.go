package cmd

import (
	"strings"
)

func GenerateAsciiArtJustify(ART, ALIGN string, width int) string {
	// Split the art into individual horizontal slices
	lines := strings.Split(ART, "\n")
	var result []string

	for _, line := range lines {
		// Calculate how much space we have to work with
		lineLen := len(line)

		// If the art is wider than the terminal, return as is or truncate
		if lineLen >= width {
			result = append(result, line)
			continue
		}

		switch ALIGN {
		case "center":
			// Distribute padding equally on both sides
			padding := (width - lineLen) / 2
			result = append(result, strings.Repeat(" ", padding)+line)

		case "left":
			// Standard behavior
			result = append(result, line)

		case "right":
			// Push to the far right
			padding := width - lineLen
			result = append(result, strings.Repeat(" ", padding)+line)

		case "justify":
			// Ensure result starts fresh for the justification process
			result = []string{}

			for i := 0; i < len(lines); i += 8 {
				end := i + 8
				if end > len(lines) {
					end = len(lines)
				}
				block := lines[i:end]

				// Skip justification if the line is empty or already too wide
				if len(block[0]) == 0 || len(block[0]) >= width {
					result = append(result, block...)
					continue
				}

				// --- Gap Detection ---
				type Gap struct {
					start int
					width int
				}
				var gaps []Gap
				currentGapStart := -1

				for col := 0; col < len(block[0]); col++ {
					isFullSpaceCol := true
					for row := 0; row < len(block); row++ {
						if col >= len(block[row]) || block[row][col] != ' ' {
							isFullSpaceCol = false
							break
						}
					}

					if isFullSpaceCol {
						if currentGapStart == -1 {
							currentGapStart = col
						}
					} else if currentGapStart != -1 {
						gapWidth := col - currentGapStart
						// Only count as a word gap if it's wider than standard letter spacing
						if gapWidth > 1 {
							gaps = append(gaps, Gap{start: currentGapStart, width: gapWidth})
						}
						currentGapStart = -1
					}
				}

				numGaps := len(gaps)
				if numGaps == 0 {
					result = append(result, block...)
					continue
				}

				// --- Calculation ---
				extraSpace := width - len(block[0])
				spacePerGap := extraSpace / numGaps
				remainder := extraSpace % numGaps

				// --- Building the New Lines ---
				for row := 0; row < len(block); row++ {
					var newLine strings.Builder
					lastCopied := 0

					for gIdx, gap := range gaps {
						// 1. Copy the art segment
						newLine.WriteString(block[row][lastCopied:gap.start])

						// 2. Add the justified gap
						padding := gap.width + spacePerGap
						if gIdx < remainder {
							padding++
						}
						newLine.WriteString(strings.Repeat(" ", padding))

						lastCopied = gap.start + gap.width
					}

					// 3. Add the final segment of the art
					if lastCopied < len(block[row]) {
						newLine.WriteString(block[row][lastCopied:])
					}
					result = append(result, newLine.String())
				}
			}

		default:
			aligns := []string{"center", "left", "right", "justify"}

			for _, align := range aligns {
				if align != ALIGN {
					return "Use center, left, right or justify\n"
				}
			}

			result = append(result, line)
		}
	}

	return strings.Join(result, "\n")
}
