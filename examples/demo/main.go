// Copyright 2024 Roi Martin.

// Demo is a simple program that shows Dear ImGui demo window.
package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/jroimartin/imgui"
	igglfw "github.com/jroimartin/imgui/backend/glfw"
	igopengl "github.com/jroimartin/imgui/backend/opengl"
)

const (
	initialWidth  = 1220
	initialHeight = 720
)

func main() {
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		log.Fatalf("failed to initialize GLFW: %v", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	win, err := glfw.CreateWindow(initialWidth, initialHeight, "Dear ImGui demo", nil, nil)
	if err != nil {
		log.Fatalf("failed to create GLFW window: %v", err)
	}
	win.MakeContextCurrent()

	win.SetFramebufferSizeCallback(framebufferSizeCB)

	if err := gl.Init(); err != nil {
		log.Fatalf("failed to initialize OpenGL")
	}

	igCtx := imgui.CreateContext(nil)
	defer igCtx.Destroy()

	if !igglfw.InitForOpenGL(win.Handle(), true) {
		log.Fatal("failed to initialize ImGui GLFW backend")
	}
	defer igglfw.Shutdown()

	if !igopengl.Init("#version 330 core") {
		log.Fatal("failed to initialize ImGui OpenGL backend")
	}
	defer igopengl.Shutdown()

	demoIsOpen := true

	for !win.ShouldClose() {
		glfw.PollEvents()

		igopengl.NewFrame()
		igglfw.NewFrame()
		imgui.NewFrame()

		if demoIsOpen {
			imgui.ShowDemoWindow(&demoIsOpen)
		}

		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		imgui.Render()
		igopengl.RenderDrawData(imgui.GetDrawData())

		win.SwapBuffers()
	}
}

func framebufferSizeCB(_win *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
}
