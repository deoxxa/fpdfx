package fpdfx

import (
	"strings"

	"github.com/jung-kurt/gofpdf"
)

func BreakLines(pdf *gofpdf.Fpdf, text string, maxWidth float64) []string {
	var lines []string

	for _, inputLine := range strings.Split(text, "\n") {
		words := strings.Split(inputLine, " ")

		var line string
		for _, word := range words {
			if pdf.GetStringWidth(strings.TrimSpace(line+" "+word)) < maxWidth {
				line = strings.TrimSpace(line + " " + word)
			} else {
				lines = append(lines, line)
				line = word
			}
		}

		lines = append(lines, strings.TrimSpace(line))
	}

	if len(lines) > 0 {
		if lines[len(lines)-1] == "" {
			lines = lines[0 : len(lines)-1]
		}
	}

	return lines
}

func MaxHeight(groups ...[]string) int {
	var highest int

	for _, e := range groups {
		if l := len(e); l > highest {
			highest = l
		}
	}

	return highest
}

func WriteLines(pdf *gofpdf.Fpdf, width, lineHeight float64, lines []string) {
	for _, line := range lines {
		pdf.CellFormat(width, lineHeight, line, "", 2, "L", false, 0, "")
	}
}
