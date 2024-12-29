// Copyright 2024 Roi Martin.

// Package imgui provides bindings for the [Dear ImGui] library.
//
// # Pre-compiled imgui static library
//
// This package bundles a pre-compiled cimgui static library to make
// it easier to use.
//
// Run the following command to build the pre-compiled cimgui static
// library for the running host as well as the corresponding header
// files:
//
//	go generate
//
// # Wayland
//
// If you are targeting wayland, specify the "wayland" build tag. For
// instance,
//
//	go run -tags wayland github.com/jroimartin/imgui/examples/glfw-opengl@latest
//
// [Dear ImGui]: https://github.com/ocornut/imgui
//
//go:generate make -f _deps/cimgui.mk
package imgui

// #cgo linux amd64 LDFLAGS: ${SRCDIR}/libcimgui_linux_amd64.a -lstdc++
// #cgo CPPFLAGS: -D CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include <stdlib.h>
// #include <stdbool.h>
// #include <cimgui.h>
import "C"

import "unsafe"

// NewFrame starts a new frame.
func NewFrame() {
	C.igNewFrame()
}

// Render renders a frame.
func Render() {
	C.igRender()
}

// GetDrawData returns the draw data required to render a frame.
func GetDrawData() *DrawData {
	return &DrawData{
		data: C.igGetDrawData(),
	}
}

// FontAtlas represents a font atlas.
type FontAtlas struct {
	data *C.ImFontAtlas
}

// NewFontAtlas returns a new [FontAtlas] value.
func NewFontAtlas() *FontAtlas {
	return &FontAtlas{
		data: C.ImFontAtlas_ImFontAtlas(),
	}
}

// Ptr returns an unsafe pointer to the underlying C type.
func (fa *FontAtlas) Ptr() unsafe.Pointer {
	return unsafe.Pointer(fa.data)
}

// GuiContext represents a Dear ImGui context.
type GuiContext struct {
	data *C.ImGuiContext
}

// CreateContext creates a context.
func CreateContext(fontAtlas *FontAtlas) *GuiContext {
	var data *C.ImFontAtlas
	if fontAtlas != nil {
		data = fontAtlas.data
	}
	return &GuiContext{
		data: C.igCreateContext(data),
	}
}

// Ptr returns an unsafe pointer to the underlying C type.
func (gc *GuiContext) Ptr() unsafe.Pointer {
	return unsafe.Pointer(gc.data)
}

// Destroy destroys the context.
func (gc *GuiContext) Destroy() {
	C.igDestroyContext(gc.data)
}

// DrawData contains the draw data required to render a frame.
type DrawData struct {
	data *C.ImDrawData
}

// NewDrawData returns a new [DrawData] value.
func NewDrawData() *DrawData {
	return &DrawData{
		data: C.ImDrawData_ImDrawData(),
	}
}

// Ptr returns an unsafe pointer to the underlying C type.
func (dd *DrawData) Ptr() unsafe.Pointer {
	return unsafe.Pointer(dd.data)
}

// ShowDemoWindow shows the Deam ImGui demo window. If open is not
// nil, it shows a window-closing widget in the upper-right corner of
// the window, which clicking will set the boolean to false when
// clicked.
func ShowDemoWindow(open *bool) {
	C.igShowDemoWindow((*C._Bool)(open))
}
