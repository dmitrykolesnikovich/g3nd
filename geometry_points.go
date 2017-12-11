package main

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
)

func init() {
	TestMap["geometry.points"] = &Points{}
}

type Points struct{}

func (t *Points) Initialize(ctx *Context) {

	// Adds axis helper to the scene
	axis := graphic.NewAxisHelper(2)
	ctx.Scene.Add(axis)

	// Creates geometry
	geom := geometry.NewGeometry()
	positions := math32.NewArrayF32(0, 0)
	positions.Append(
		1, 1, 1,
		1, 1, -1,
		-1, 1, -1,
		-1, 1, 1,
		1, -1, 1,
		1, -1, -1,
		-1, -1, -1,
		-1, -1, 1,
		0, 1, 0,
		0, 0, -1,
		0, -1, 0,
		1, 0, 0,
		-1, 0, 0,
		0, 0, 1,
		0, 0, 0,
	)
	geom.AddVBO(gls.NewVBO().AddAttrib("VertexPosition", 3).SetBuffer(positions))
	positions = nil // Positions cannot be used after transfering to VBO

	// Creates point material
	mat := material.NewPoint(&math32.Color{0, 0, 0})
	mat.SetSize(2000)

	// Creates points mesh
	mesh := graphic.NewPoints(geom, mat)
	mesh.SetScale(1, 1, 1)
	ctx.Scene.Add(mesh)
}

func (t *Points) Render(ctx *Context) {
}
