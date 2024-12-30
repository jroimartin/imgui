// Copyright 2024 Roi Martin.

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

// WindowFlags represents the window flags.
type WindowFlags C.ImGuiWindowFlags

// Window flags.
const (
	WindowFlagsAlwaysAutoresize = WindowFlags(C.ImGuiWindowFlags_AlwaysAutoResize)
)

// ColorEditFlags represents the color edit flags.
type ColorEditFlags C.ImGuiColorEditFlags

// Color edit flags.
const (
	ColorEditFlagsNoInputs = ColorEditFlags(C.ImGuiColorEditFlags_NoInputs)
)

// Begin pushes a new window to the stack to start appending widgets
// to it. If open is not nil, it shows a window-closing widget in the
// upper-right corner of the window, which clicking will set the
// boolean to false when clicked. It returns false if the window is
// collapsed.
func Begin(name string, open *bool, flags WindowFlags) bool {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cOpen := (*C._Bool)(open)
	cFlags := C.ImGuiWindowFlags(flags)
	return bool(C.igBegin(cName, cOpen, cFlags))
}

// End pops a window from the stack.
func End() {
	C.igEnd()
}

// SetNextWindowPos sets the positions of the next window.
func SetNextWindowPos(pos Vec2, cond int, pivot Vec2) {
	cPos := pos.toImVec2()
	cCond := C.ImGuiCond(cond)
	cPivot := pivot.toImVec2()
	C.igSetNextWindowPos(cPos, cCond, cPivot)
}

// SetNextWindowSize sets the size of the next window.
func SetNextWindowSize(size Vec2, cond int) {
	cSize := size.toImVec2()
	cCond := C.ImGuiCond(cond)
	C.igSetNextWindowSize(cSize, cCond)
}

// ColorEdit4 adds a color picker widget. col reports the selected
// color. It returns whether the color has changed.
func ColorEdit4(label string, col *[4]float32, flags ColorEditFlags) bool {
	cLabel := C.CString(label)
	defer C.free(unsafe.Pointer(cLabel))
	cCol := (*C.float)(&col[0])
	cFlags := C.ImGuiColorEditFlags(flags)
	retval := C.igColorEdit4(cLabel, cCol, cFlags)
	return bool(retval)
}

// ShowDemoWindow shows the Deam ImGui demo window. If open is not
// nil, it shows a window-closing widget in the upper-right corner of
// the window, which clicking will set the boolean to false when
// clicked.
func ShowDemoWindow(open *bool) {
	C.igShowDemoWindow((*C._Bool)(open))
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
func (ctx *GuiContext) Ptr() unsafe.Pointer {
	return unsafe.Pointer(ctx.data)
}

// Destroy destroys the context.
func (ctx *GuiContext) Destroy() {
	C.igDestroyContext(ctx.data)
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
func (data *DrawData) Ptr() unsafe.Pointer {
	return unsafe.Pointer(data.data)
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
func (atlas *FontAtlas) Ptr() unsafe.Pointer {
	return unsafe.Pointer(atlas.data)
}

// IO represents the IO state.
type IO struct {
	data *C.ImGuiIO
}

// GetIO returns the IO state.
func GetIO() *IO {
	return &IO{data: C.igGetIO()}
}

// Ptr returns an unsafe pointer to the underlying C type.
func (io *IO) Ptr() unsafe.Pointer {
	return unsafe.Pointer(io.data)
}

// ConfigFlags represents the configuration flags.
type ConfigFlags = C.ImGuiConfigFlags

// Configuration flags.
const (
	ConfigFlagsNavEnableKeyboard = ConfigFlags(C.ImGuiConfigFlags_NavEnableKeyboard)
	ConfigFlagsDockingEnable     = ConfigFlags(C.ImGuiConfigFlags_DockingEnable)
)

// SetConfigFlags sets the configuration flags.
func (io *IO) SetConfigFlags(flags ConfigFlags) {
	io.data.ConfigFlags = C.ImGuiConfigFlags(flags)
}

// ConfigFlags returns the configuration flags.
func (io *IO) ConfigFlags() ConfigFlags {
	return ConfigFlags(io.data.ConfigFlags)
}

// SetIniFilename sets the path of the .ini file. If name is empty, it
// disables automatic load/save. Note that this function allocates a C
// string internally which is leaked.
func (io *IO) SetIniFilename(name string) {
	var cname *C.char
	if name != "" {
		cname = C.CString(name)
	}
	io.data.IniFilename = cname
}

// SetLogFilename sets the path of the .log file. If name is empty, it
// disables logging. Note that this function allocates a C string
// internally which is leaked.
func (io *IO) SetLogFilename(name string) {
	var cname *C.char
	if name != "" {
		cname = C.CString(name)
	}
	io.data.LogFilename = cname
}

// Viewport represents the platform Window created by the application
// which is hosting the Dear ImGui windows.
type Viewport struct {
	data *C.ImGuiViewport
}

// GetMainViewport returns the primary/default viewport.
func GetMainViewport() *Viewport {
	return &Viewport{data: C.igGetMainViewport()}
}

// Ptr returns an unsafe pointer to the underlying C type.
func (vp *Viewport) Ptr() unsafe.Pointer {
	return unsafe.Pointer(vp.data)
}

// GetWorkpos returns the position of the viewport minus task bars,
// menus bars and status bars.
func (vp *Viewport) GetWorkpos() Vec2 {
	return newVec2FromImVec2(vp.data.WorkPos)
}

// GetWorksize returns the size of the viewport minus task bars, menus
// bars and status bars.
func (vp *Viewport) GetWorksize() Vec2 {
	return newVec2FromImVec2(vp.data.WorkSize)
}

// Vec2 represents a two-dimensional vector.
type Vec2 struct {
	X, Y float32
}

// fromImVec2 creates a new [Vec2] from its C counterpart.
func newVec2FromImVec2(v C.ImVec2) Vec2 {
	return Vec2{X: float32(v.x), Y: float32(v.y)}
}

// toImVec2 converts a [Vec2] into its C counterpart.
func (v *Vec2) toImVec2() C.ImVec2 {
	return C.ImVec2{x: C.float(v.X), y: C.float(v.Y)}
}
