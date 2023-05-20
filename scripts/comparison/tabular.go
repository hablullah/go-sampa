package main

import "strings"

type Table struct {
	Rows [][]string
}

func (t *Table) AddRow(cells ...string) {
	// If there is an existing row, make sure column count match
	if len(t.Rows) > 0 {
		colCount := len(t.Rows[0])
		if len(cells) != colCount {
			return
		}
	}

	t.Rows = append(t.Rows, cells)
}

func (t Table) String() string {
	if len(t.Rows) == 0 {
		return ""
	}

	// Calculate column width
	colWidths := make([]int, len(t.Rows[0]))
	for _, row := range t.Rows {
		for col, cell := range row {
			width := len(cell) + 2
			currentWidth := colWidths[col]
			if width > currentWidth {
				colWidths[col] = width
			}
		}
	}

	// Build strings
	var sb strings.Builder
	sb.WriteString(t.rowString(t.Rows[0], colWidths))
	sb.WriteString("\n")
	sb.WriteString(t.headerSeparator(colWidths))

	for i := 1; i < len(t.Rows); i++ {
		sb.WriteString("\n")
		sb.WriteString(t.rowString(t.Rows[i], colWidths))
	}

	return sb.String()
}

func (t Table) spacePad(s string, width int) string {
	sWidth := len(s)
	lWidth := (width - sWidth) / 2
	rWidth := width - lWidth - sWidth
	return strings.Repeat(" ", rWidth) + s + strings.Repeat(" ", lWidth)
}

func (t Table) rowString(cells []string, colWidths []int) string {
	for col, cell := range cells {
		cells[col] = t.spacePad(cell, colWidths[col])
	}
	return "|" + strings.Join(cells, "|") + "|"
}

func (t Table) headerSeparator(colWidths []int) string {
	cells := make([]string, len(colWidths))
	for i, w := range colWidths {
		cells[i] = ":" + strings.Repeat("-", w-2) + ":"
	}
	return "|" + strings.Join(cells, "|") + "|"
}
