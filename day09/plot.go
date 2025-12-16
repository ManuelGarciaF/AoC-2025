package main

import (
	"github.com/ManuelGarciaF/AoC-2025/commons"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func plotBox(title string, x1, x2 commons.Coord, coords []commons.Coord) {
	p := plot.New()

	p.Title.Text = title

	plotutil.AddLinePoints(p,
		"Points", plotter.XYs(commons.Map(coords, func(c commons.Coord) plotter.XY {
			return plotter.XY{
				X: float64(c.X),
				Y: float64(c.Y),
			}
		})),
		"Solution", plotter.XYs{
			{X: float64(x1.X),  Y: float64(x1.Y)},
			{X: float64(x1.X),  Y: float64(x2.Y)},
			{X: float64(x2.X),  Y: float64(x2.Y)},
			{X: float64(x2.X),  Y: float64(x1.Y)},
			{X: float64(x1.X),  Y: float64(x1.Y)},
		},
	)

	p.Save(30*vg.Centimeter, 30*vg.Centimeter, "imgs/" + title + ".png")
}
