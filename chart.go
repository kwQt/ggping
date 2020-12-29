package ggping

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func DrawChart(plot *widgets.Plot, data []float64, max float64, width int, height int) {
	if len(data) < 2 {
		return
	}

	plot.Marker = widgets.MarkerBraille
	plot.Data = [][]float64{data}
	plot.SetRect(0, 0, width, height)
	plot.MaxVal = max
	plot.DotMarkerRune = '+'
	plot.AxesColor = ui.ColorWhite
	plot.LineColors[0] = ui.ColorYellow
	plot.DrawDirection = widgets.DrawLeft
	plot.Border = false

	ui.Render(plot)
}
