// Copyright 2024 Roi Martin.

// Package opengl provides a Dear ImGui OpenGL backend.
package opengl

// #cgo CPPFLAGS: -I ${SRCDIR}/../..
// #cgo CPPFLAGS: -D CIMGUI_DEFINE_ENUMS_AND_STRUCTS -D CIMGUI_USE_OPENGL3
// #include <stdlib.h>
// #include <cimgui.h>
// #include <cimgui_impl.h>
import "C"

import (
	"unsafe"

	"github.com/jroimartin/imgui"
)

// Init initializes the OpenGL backend.
func Init(glslVersion string) bool {
	cGlslVersion := C.CString(glslVersion)
	defer C.free(unsafe.Pointer(cGlslVersion))
	return bool(C.ImGui_ImplOpenGL3_Init(cGlslVersion))
}

// Shutdown terminates the OpenGL backend.
func Shutdown() {
	C.ImGui_ImplOpenGL3_Shutdown()
}

// NewFrame starts a frame.
func NewFrame() {
	C.ImGui_ImplOpenGL3_NewFrame()
}

// RenderDrawData renders draw data.
func RenderDrawData(drawData *imgui.DrawData) {
	cDrawData := (*C.ImDrawData)(drawData.Ptr())
	C.ImGui_ImplOpenGL3_RenderDrawData(cDrawData)
}
