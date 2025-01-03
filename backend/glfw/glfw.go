// Copyright 2024 Roi Martin.

// Package glfw provides a Dear ImGui GLFW backend.
package glfw

// #cgo CPPFLAGS: -I ${SRCDIR}/../..
// #cgo CPPFLAGS: -D CIMGUI_DEFINE_ENUMS_AND_STRUCTS -D CIMGUI_USE_GLFW
// #include <cimgui.h>
// #include <cimgui_impl.h>
import "C"
import "unsafe"

// InitForOpenGL initializes the GLFW backend for OpenGL. win is a
// GLFWwindow C pointer.
func InitForOpenGL(win unsafe.Pointer, installCallbacks bool) bool {
	cWin := (*C.GLFWwindow)(win)
	cInstallCallbacks := C._Bool(installCallbacks)
	return bool(C.ImGui_ImplGlfw_InitForOpenGL(cWin, cInstallCallbacks))
}

// Shutdown terminates the GLFW backend.
func Shutdown() {
	C.ImGui_ImplGlfw_Shutdown()
}

// NewFrame starts a new frame.
func NewFrame() {
	C.ImGui_ImplGlfw_NewFrame()
}
